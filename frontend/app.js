let app = new Vue({
    el:"#app",
    data:{
        test:"success",
        selectedChat:-1,
        chats:[],
        selectedChatMessage:[],
        
    },
    methods:{
        addChat:function(){
            this.chats.push({Name:"test",ChatHistory:[]});
        },
        // addMessage: function(index, message){
        //     this.chats[index].ChatHistory.push(message);
        // },
        addMessage: function( message){
            this.chats[this.selectedChat].ChatHistory.push(message);
        },

        changeChat:function(index){
            this.selectedChatMessage = this.chats[index].ChatHistory;
            this.selectedChat = index;
        }

    }
})