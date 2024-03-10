
declare module 'api' {
    interface MyApi{
      getProblemList(param: any):Promise<any>;
      getProblemDetail(param: any):Promise<any>;
      getSortList(param: any):Promise<any>;
      getRankList(param: any):Promise<any>;
      getSubmitList(param: any):Promise<any>;
      sendCode(param: any):Promise<any>;
      login(param: any):Promise<any>;
      register(param: any):Promise<any>;
      delSort(param: any):Promise<any>;
      addSort(param: any):Promise<any>;
      addProblem(param: any):Promise<any>;
      editProblem(param: any):Promise<any>;
      editSort(param: any):Promise<any>;
      submitCode(param: any, id:any):Promise<any>;
      uploadFile(param: any):Promise<any>;
    }
  
    const api: MyApi;
    export default api;
}

/*declare function getProblemList(params:any): Promise<any>;
declare function login(params:any): Promise<any>;
declare function getProblemDetail(params:any): Promise<any>;*/
