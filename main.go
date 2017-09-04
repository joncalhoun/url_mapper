package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var mappings map[string]string
	err := json.Unmarshal([]byte(data), &mappings)
	if err != nil {
		panic(err)
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		if dest, ok := mappings[r.URL.Path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fmt.Println(r.URL.Path)
		http.Error(w, "Page not found", http.StatusNotFound)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

var data = `
{
  "/2.1": "https://github.com/joncalhoun/lenslocked.com/tree/book-2.1-Building_the_server",
  "/3.1": "https://github.com/joncalhoun/lenslocked.com/tree/book-3.1-Routing_with_if_else_statements",
  "/3.1-diff": "https://github.com/joncalhoun/lenslocked.com/compare/book-2.1-Building_the_server...book-3.1-Routing_with_if_else_statements",
  "/3.3": "https://github.com/joncalhoun/lenslocked.com/tree/book-3.3-Using_the_gorilla_mux_router",
  "/3.3-diff": "https://github.com/joncalhoun/lenslocked.com/compare/book-3.1-Routing_with_if_else_statements...book-3.3-Using_the_gorilla_mux_router",
  "/4.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-4.4-Creating_a_template",
  "/4.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-3.3-Using_the_gorilla_mux_router...book-4.4-Creating_a_template",
  "/4.5": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-4.5-Contextual_encoding",
  "/4.5-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-4.4-Creating_a_template...book-4.5-Contextual_encoding",
  "/6.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.1-The_home_template",
  "/6.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-4.5-Contextual_encoding...book-6.1-The_home_template",
  "/6.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.2-The_contact_template",
  "/6.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.1-The_home_template...book-6.2-The_contact_template",
  "/6.3.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.3.1-Footer_template",
  "/6.3.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.2-The_contact_template...book-6.3.1-Footer_template",
  "/6.3.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.3.2-Creating_the_view_type",
  "/6.3.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.3.1-Footer_template...book-6.3.2-Creating_the_view_type",
  "/6.3.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.3.3-Bootstrap_layout",
  "/6.3.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.3.2-Creating_the_view_type...book-6.3.3-Bootstrap_layout",
  "/6.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.4-Navbar",
  "/6.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.3.3-Bootstrap_layout...book-6.4-Navbar",
  "/6.5.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.5.2-Globbing",
  "/6.5.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.4-Navbar...book-6.5.2-Globbing",
  "/6.5.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.5.3-View_rendering",
  "/6.5.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.5.2-Globbing...book-6.5.3-View_rendering",
  "/6.5.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-6.5.4-Move_footer",
  "/6.5.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.5.3-View_rendering...book-6.5.4-Move_footer",
  "/7.1.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.1.1-Sign_up_form",
  "/7.1.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-6.5.4-Move_footer...book-7.1.1-Sign_up_form",
  "/7.1.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.1.2-Form_panel",
  "/7.1.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.1.1-Sign_up_form...book-7.1.2-Form_panel",
  "/7.1.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.1.3-Sign_up_link",
  "/7.1.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.1.2-Form_panel...book-7.1.3-Sign_up_link",
  "/7.3.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.3.1-Users_controller",
  "/7.3.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.1.3-Sign_up_link...book-7.3.1-Users_controller",
  "/7.3.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.3.2-Mv_signup",
  "/7.3.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.3.1-Users_controller...book-7.3.2-Mv_signup",
  "/7.3.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.3.3-Using_controllers",
  "/7.3.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.3.2-Mv_signup...book-7.3.3-Using_controllers",
  "/7.4.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.4.1-User_create_stub",
  "/7.4.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.3.3-Using_controllers...book-7.4.1-User_create_stub",
  "/7.4.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.4.2-Router_methods",
  "/7.4.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.4.1-User_create_stub...book-7.4.2-Router_methods",
  "/7.4.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.4.3-Parse_form",
  "/7.4.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.4.2-Router_methods...book-7.4.3-Parse_form",
  "/7.4.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.4.4-Gorilla_schema",
  "/7.4.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.4.3-Parse_form...book-7.4.4-Gorilla_schema",
  "/7.4.5": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.4.5-DRY",
  "/7.4.5-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.4.4-Gorilla_schema...book-7.4.5-DRY",
  "/7.5.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.5.1-Static_controller",
  "/7.5.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.4.5-DRY...book-7.5.1-Static_controller",
  "/7.5.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-7.5.2-Simpler_views",
  "/7.5.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.5.1-Static_controller...book-7.5.2-Simpler_views",
  "/8.3.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-8.3.1-PostgreSQL",
  "/8.3.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-7.5.2-Simpler_views...book-8.3.1-PostgreSQL",
  "/8.3.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-8.3.3-Writing_sql_records",
  "/8.3.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-8.3.1-PostgreSQL...book-8.3.3-Writing_sql_records",
  "/8.3.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-8.3.4-Query_single_record",
  "/8.3.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-8.3.3-Writing_sql_records...book-8.3.4-Query_single_record",
  "/8.3.5": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-8.3.5-Multiple_records",
  "/8.3.5-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-8.3.4-Query_single_record...book-8.3.5-Multiple_records",
  "/8.3.6": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-8.3.6-Writing_relational_records",
  "/8.3.6-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-8.3.5-Multiple_records...book-8.3.6-Writing_relational_records",
  "/8.3.7": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-8.3.7-Query_related_records",
  "/8.3.7-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-8.3.6-Writing_relational_records...book-8.3.7-Query_related_records",
  "/8.4.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-8.4.2-GORM_model",
  "/8.4.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-8.3.7-Query_related_records...book-8.4.2-GORM_model",
  "/8.4.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-8.4.3-Automigrate",
  "/8.4.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-8.4.2-GORM_model...book-8.4.3-Automigrate",
  "/8.4.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-8.4.4-Logging",
  "/8.4.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-8.4.3-Automigrate...book-8.4.4-Logging",
  "/8.4.5": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-8.4.5-Creating_records",
  "/8.4.5-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-8.4.4-Logging...book-8.4.5-Creating_records",
  "/8.4.6": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-8.4.6-Query_single",
  "/8.4.6-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-8.4.5-Creating_records...book-8.4.6-Query_single",
  "/8.4.7": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-8.4.7-Query_multi",
  "/8.4.7-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-8.4.6-Query_single...book-8.4.7-Query_multi",
  "/8.4.8": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-8.4.8-Create_GORM_relational",
  "/8.4.8-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-8.4.7-Query_multi...book-8.4.8-Create_GORM_relational",
  "/8.4.9": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-8.4.9-Query_GORM_relational",
  "/8.4.9-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-8.4.8-Create_GORM_relational...book-8.4.9-Query_GORM_relational",
  "/9.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-9.1-Defining_user",
  "/9.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-8.4.9-Query_GORM_relational...book-9.1-Defining_user",
  "/9.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-9.2-UserService",
  "/9.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-9.1-Defining_user...book-9.2-UserService",
  "/9.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-9.3-Create_user",
  "/9.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-9.2-UserService...book-9.3-Create_user",
  "/9.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-9.4-ByEmail_and_DRY",
  "/9.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-9.3-Create_user...book-9.4-ByEmail_and_DRY",
  "/9.5": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-9.5-Update_delete",
  "/9.5-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-9.4-ByEmail_and_DRY...book-9.5-Update_delete",
  "/9.6": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-9.6-automigrate",
  "/9.6-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-9.5-Update_delete...book-9.6-automigrate",
  "/9.7.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-9.7.1-Form_name",
  "/9.7.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-9.6-automigrate...book-9.7.1-Form_name",
  "/9.7.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-9.7.2-Setup_UserService",
  "/9.7.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-9.7.1-Form_name...book-9.7.2-Setup_UserService",
  "/9.7.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-9.7.3-UserService_in_controller",
  "/9.7.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-9.7.2-Setup_UserService...book-9.7.3-UserService_in_controller",
  "/10.4.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-10.4.1-UserPwField",
  "/10.4.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-9.7.3-UserService_in_controller...book-10.4.1-UserPwField",
  "/10.4.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-10.4.2-HashPassword",
  "/10.4.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-10.4.1-UserPwField...book-10.4.2-HashPassword",
  "/10.4.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-10.4.3-UsePwInController",
  "/10.4.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-10.4.2-HashPassword...book-10.4.3-UsePwInController",
  "/10.4.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-10.4.4-SaltAndPepper",
  "/10.4.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-10.4.3-UsePwInController...book-10.4.4-SaltAndPepper",
  "/10.5.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-10.5.1-LoginTmpl",
  "/10.5.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-10.4.4-SaltAndPepper...book-10.5.1-LoginTmpl",
  "/10.5.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-10.5.2-LoginAction",
  "/10.5.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-10.5.1-LoginTmpl...book-10.5.2-LoginAction",
  "/10.5.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-10.5.3-AuthImpl",
  "/10.5.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-10.5.2-LoginAction...book-10.5.3-AuthImpl",
  "/10.5.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-10.5.4-CallAuth",
  "/10.5.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-10.5.3-AuthImpl...book-10.5.4-CallAuth",
  "/11.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-11.2-First_cookie",
  "/11.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-10.5.4-CallAuth...book-11.2-First_cookie",
  "/11.3.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-11.3.2-Read_cookie",
  "/11.3.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-11.2-First_cookie...book-11.3.2-Read_cookie",
  "/11.5": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-11.5-Remember_tokens",
  "/11.5-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-11.3.2-Read_cookie...book-11.5-Remember_tokens",
  "/11.6.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-11.6.3-hash_pkg",
  "/11.6.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-11.5-Remember_tokens...book-11.6.3-hash_pkg",
  "/11.7.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-11.7.1-Remember_fields",
  "/11.7.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-11.6.3-hash_pkg...book-11.7.1-Remember_fields",
  "/11.7.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-11.7.2-Set_remember_on_create",
  "/11.7.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-11.7.1-Remember_fields...book-11.7.2-Set_remember_on_create",
  "/11.7.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-11.7.3-HMAC_in_US",
  "/11.7.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-11.7.2-Set_remember_on_create...book-11.7.3-HMAC_in_US",
  "/11.7.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-11.7.4-Hash_tokens",
  "/11.7.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-11.7.3-HMAC_in_US...book-11.7.4-Hash_tokens",
  "/11.7.5": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-11.7.5-ByRemember",
  "/11.7.5-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-11.7.4-Hash_tokens...book-11.7.5-ByRemember",
  "/11.7.6": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-11.7.6-Test_it",
  "/11.7.6-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-11.7.5-ByRemember...book-11.7.6-Test_it",
  "/11.8.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-11.8.1-Token_cookies",
  "/11.8.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-11.7.6-Test_it...book-11.8.1-Token_cookies",
  "/11.8.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-11.8.2-Restricting_access",
  "/11.8.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-11.8.1-Token_cookies...book-11.8.2-Restricting_access",
  "/11.9": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-11.9-XSS",
  "/11.9-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-11.8.2-Restricting_access...book-11.9-XSS",
  "/12.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.2-UserDB",
  "/12.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-11.9-XSS...book-12.2-UserDB",
  "/12.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.3-UserService_interface",
  "/12.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.2-UserDB...book-12.3-UserService_interface",
  "/12.4.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.4.1-Val_example",
  "/12.4.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.3-UserService_interface...book-12.4.1-Val_example",
  "/12.4.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.4.2-More_val_ex",
  "/12.4.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.4.1-Val_example...book-12.4.2-More_val_ex",
  "/12.4.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.4.3-Reusable_vals",
  "/12.4.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.4.2-More_val_ex...book-12.4.3-Reusable_vals",
  "/12.5.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.5.1-Remember_token_norm",
  "/12.5.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.4.3-Reusable_vals...book-12.5.1-Remember_token_norm",
  "/12.5.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.5.2-Default_remember",
  "/12.5.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.5.1-Remember_token_norm...book-12.5.2-Default_remember",
  "/12.5.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.5.3-Min_id",
  "/12.5.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.5.2-Default_remember...book-12.5.3-Min_id",
  "/12.6.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.6.1-Email_norm",
  "/12.6.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.5.3-Min_id...book-12.6.1-Email_norm",
  "/12.6.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.6.2-Req_email",
  "/12.6.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.6.1-Email_norm...book-12.6.2-Req_email",
  "/12.6.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.6.3-Email_regex",
  "/12.6.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.6.2-Req_email...book-12.6.3-Email_regex",
  "/12.6.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.6.4-Email_avail",
  "/12.6.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.6.3-Email_regex...book-12.6.4-Email_avail",
  "/12.7": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.7-Err_rename",
  "/12.7-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.6.4-Email_avail...book-12.7-Err_rename",
  "/12.8.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.8.1-Pw_len",
  "/12.8.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.7-Err_rename...book-12.8.1-Pw_len",
  "/12.8.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.8.2-Pw_req",
  "/12.8.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.8.1-Pw_len...book-12.8.2-Pw_req",
  "/12.9": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-12.9-Rem_vals",
  "/12.9-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.8.2-Pw_req...book-12.9-Rem_vals",
  "/13.1": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-13.1-Bootstrap_alert",
  "/13.1-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-12.9-Rem_vals...book-13.1-Bootstrap_alert",
  "/13.2": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-13.2-Dynamic_alerts",
  "/13.2-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-13.1-Bootstrap_alert...book-13.2-Dynamic_alerts",
  "/13.3": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-13.3-Conditional_alerts",
  "/13.3-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-13.2-Dynamic_alerts...book-13.3-Conditional_alerts",
  "/13.4": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-13.4-View_data",
  "/13.4-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-13.3-Conditional_alerts...book-13.4-View_data",
  "/13.5": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-13.5-Signup_err",
  "/13.5-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-13.4-View_data...book-13.5-Signup_err",
  "/13.6": "https://gitlab.com/joncalhoun/lenslocked.com/tree/book-13.6-Public_errors",
  "/13.6-diff": "https://gitlab.com/joncalhoun/lenslocked.com/compare/book-13.5-Signup_err...book-13.6-Public_errors"
}
`

// env GOOS=linux GOARCH=amd64 go build . && ssh root@members.usegolang.com "sudo service url_mapper stop" && scp url_mapper root@members.usegolang.com:~/url_mapper && ssh root@members.usegolang.com "sudo service url_mapper restart"
