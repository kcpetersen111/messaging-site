let app = new Vue({
    el:"#app",
    data:{
        test:"success",
        selectedChat:-1,
        username:"new user",
        chats:[],
        selectedChatMessage:[],
        message:"",
        showChatSearch:false,
        chatSearchUser:"",
        addUserToChat:"",
        showUserSearch:-1,
        connection:null,
        
    },
    methods:{
        addChat:function(){
            this.showChatSearch = true;
            // this.chats.push({Name:"test",ChatHistory:[]});
        },
        createChat:function(){
            this.chats.push({Name:this.chatSearchUser,ChatHistory:[], Participants:[this.username]});
            this.showChatSearch = false;
            this.chatSearchUser = "";

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
        },
        addParticipant: function(index){
            this.chats[index].Participants.push(this.addUserToChat);
            this.chats[index].ChatHistory.push({Message:`${this.addUserToChat} has joined the chat`,
             Sender:""});
            this.addUserToChat = "";
            this.showUserSearch = -1;

        } 

    },
    created: function(){
        console.log("opening a websocket");
        this.connection = new WebSocket("ws://localhost:8080");

        this.connection.onopen = (event) =>{
            console.log("websocket opened successfully");
            console.log(event);
        }
    }
})