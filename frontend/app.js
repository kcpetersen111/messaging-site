let app = new Vue({
    el:"#app",
    data:{
        test:"success",
        selectedChat:-1,
        username:"new user",
        chats:[],
        selectedChatMessage:[],
        message:"",
        
    },
    methods:{
        addChat:function(){
            this.chats.push({Name:"test",ChatHistory:[]});
        },
        // addMessage: function(index, message){
        //     this.chats[index].ChatHistory.push(message);
        // },
        addMessage: function( ){
            this.chats[this.selectedChat].ChatHistory.push({Message:this.message, Sender:this.username});
            this.message = ""
        },

        changeChat:function(index){
            this.selectedChatMessage = this.chats[index].ChatHistory;
            this.selectedChat = index;
        }

    }
})