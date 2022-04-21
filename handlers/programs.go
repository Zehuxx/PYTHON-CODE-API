package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/zehuxx/python-code-api/db"
	"github.com/zehuxx/python-code-api/helpers"
	"github.com/zehuxx/python-code-api/models"
)

//GetPrograms return list of programs data
func GetPrograms(res http.ResponseWriter, req *http.Request) {
	dg, cancel := db.GetDgraphClient()
	defer cancel()

	ctx, toCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer toCancel()

	txn := dg.NewReadOnlyTxn()

	q := `{
			programs(func: has(programName)){
				programName
				uid
			}
		}`

	apiRes, err := txn.Query(ctx, q)
	if err != nil {
		fmt.Println("Failed to get programs.", err)
		//http.Error(res, "Failed to get programs.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to get programs."}, http.StatusInternalServerError)
		return
	}

	var result models.Data
	if err := json.Unmarshal(apiRes.Json, &result); err != nil {
		fmt.Println("Failed to Unmarshal JSON.", err)
		//http.Error(res, "Failed to get programs.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to get programs."}, http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

//GetProgramByUid return program data filter by Uid
func GetProgramByUid(res http.ResponseWriter, req *http.Request) {
	dg, cancel := db.GetDgraphClient()
	defer cancel()

	ctx, toCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer toCancel()

	txn := dg.NewReadOnlyTxn()

	uid := req.Context().Value("uid").(string)

	vars := make(map[string]string)
	vars["$uid"] = uid

	q := `query q($uid: string){
			programs(func: uid($uid)) @recurse{
				programName
				id
				nid
				name
				data
				nodes(orderasc: id)
				class
				typenode
				html
				inputs
				node
				input 
				pos_x
				pos_y
				input_1  
				input_2  
				input_3  
				connections(orderasc: order)
				output
				outputs 
				output_1
				output_2
				output_3
				con
				num
				msg
				assign
			}
		}`

	apiRes, err := txn.QueryWithVars(ctx, q, vars)
	if err != nil {
		fmt.Println("Failed to get program.", err)
		//http.Error(res, "Failed to get program.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to get program."}, http.StatusInternalServerError)
		return
	}

	var result models.Data
	if err := json.Unmarshal(apiRes.Json, &result); err != nil {
		fmt.Println("Failed to Unmarshal.", err)
		//http.Error(res, "Failed to get program.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to get program."}, http.StatusInternalServerError)
		return
	}

	if len(result.Programs) == 0 {
		fmt.Println("The program was not found.", err)
		//http.Error(res, "Failed to get program.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "The program was not found."}, http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

//SaveProgram save program data
func SaveProgram(res http.ResponseWriter, req *http.Request) {
	dg, cancel := db.GetDgraphClient()
	defer cancel()

	ctx, toCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer toCancel()

	txn := dg.NewTxn()

	var program models.Programs
	// JSON to struct
	if err := json.NewDecoder(req.Body).Decode(&program); err != nil {
		fmt.Println("Failed to Unmarshal JSON.", err)
		//http.Error(res, "Failed to save program.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to save program."}, http.StatusInternalServerError)
		return
	}

	// struct to JSON
	pb, err := json.Marshal(program)
	if err != nil {
		fmt.Println("Failed to Marshal.", err)
		//http.Error(res, "Failed to save program.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to save program."}, http.StatusInternalServerError)
		return
	}

	mu := &api.Mutation{
		SetJson: pb,
	}

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		fmt.Println("Failed to set mutation.", err)
		//http.Error(res, "Failed to save program.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to save program."}, http.StatusInternalServerError)
		return
	}

	defer txn.Discard(ctx)

	err = txn.Commit(ctx)
	if err != nil {
		fmt.Println("Failed to commit.", err)
		//http.Error(res, "Failed to update program.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to save program."}, http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
}

//UpdateProgram update program data
func UpdateProgram(res http.ResponseWriter, req *http.Request) {
	dg, cancel := db.GetDgraphClient()
	defer cancel()

	ctx, toCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer toCancel()

	txn := dg.NewTxn()

	uid := req.Context().Value("uid").(string)

	var program models.Programs
	// JSON to struct
	if err := json.NewDecoder(req.Body).Decode(&program); err != nil {
		fmt.Println("Failed to Unmarshal JSON.", err)
		//http.Error(res, "Failed to update program.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to update program."}, http.StatusInternalServerError)
		return
	}

	// struct to JSON
	pb, err := json.Marshal(program)
	if err != nil {
		fmt.Println("Failed to Marshal.", err)
		//http.Error(res, "Failed to update program.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to update program."}, http.StatusInternalServerError)
		return
	}

	mu := &api.Mutation{}

	t := fmt.Sprintf("<%s> <nodes> * .", uid)
	mu.DelNquads = []byte(t)

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		fmt.Println("Failed to delete mutation.", err)
		//http.Error(res, "Failed to update program.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to update program."}, http.StatusInternalServerError)
		return
	}

	mu = &api.Mutation{}
	mu.SetJson = pb

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		fmt.Println("Failed to set mutation.", err)
		//http.Error(res, "Failed to update program.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to update program."}, http.StatusInternalServerError)
		return
	}

	defer txn.Discard(ctx)

	err = txn.Commit(ctx)
	if err != nil {
		fmt.Println("Failed to commit.", err)
		//http.Error(res, "Failed to update program.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to update program."}, http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusNoContent)
}

//DeleteProgram delete program
func DeleteProgram(res http.ResponseWriter, req *http.Request) {
	dg, cancel := db.GetDgraphClient()
	defer cancel()

	ctx, toCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer toCancel()

	txn := dg.NewTxn()

	uid := req.Context().Value("uid").(string)

	mu := &api.Mutation{}

	t := fmt.Sprintf("<%s> <nodes> * .\n <%s> <programName> * .", uid, uid)
	mu.DelNquads = []byte(t)

	_, err := txn.Mutate(ctx, mu)
	if err != nil {
		fmt.Println("Failed to delete mutation.", err)
		//http.Error(res, "Failed to delete program.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to delete program."}, http.StatusInternalServerError)
		return
	}

	defer txn.Discard(ctx)

	err = txn.Commit(ctx)
	if err != nil {
		fmt.Println("Failed to commit.", err)
		//http.Error(res, "Failed to update program.", http.StatusInternalServerError)
		helpers.JSONError(res, helpers.ErrorResponse{Msg: "Failed to delete program."}, http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusNoContent)

}
