//<reference types="vite/client" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/ban-types
  const component: DefineComponent<{}, {}, any>
  export default component
}

declare module 'vuex'

/*
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
*/
//declare module 'api/api.js'
