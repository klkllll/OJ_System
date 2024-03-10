import {createStore} from 'vuex'
const store = createStore({
    state:{
        collapse:false,
        isLogin:false,
        token:null,
        username:null,
        is_admin:0
    },
    mutations:{ 
        changeCollapse(state: { collapse: any },data: any){
                console.log(data)
                state.collapse=data
        },
        loginSucc(state: { token: any; isLogin: boolean },data: any){
            state.token=data
            state.isLogin=true
            
    },
    setUser(state: { username: any; is_admin: any },data: { username: any; is_admin: any }){
        state.username=data.username
        state.is_admin=data.is_admin
    },
    logout(state: { isLogin: boolean },data: any){
        state.isLogin=false

    }
    }

})
export default store
/* 这是之前的写法
mutations:{
    changeCollapse(state,data){
            console.log(data)
            state.collapse=data
    },
    loginSucc(state,data){
        state.token=data
        state.isLogin=true
        
},
setUser(state,data){
    state.username=data.username
    state.is_admin=data.is_admin
},
logout(state,data){
    state.isLogin=false

}
}*/