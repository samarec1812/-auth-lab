const EventHandling = {
    data() {
        return {
            message: 'Привет, Vue.js!',
            usering: [
                "admin",
                "Alexey",
                "Alyona",
                "Raf"
            ]
        }
    },
    methods: {
        //     getUsers() {
        //     // GET /someUrl
        //     this.$http.get('http://localhost:7003/auth/home/aux').then(response => {
        //         console.log(response)
        //         // get body data
        //        //  this.someData = response.body;
        //
        //     }, response => {
        //        console.log("error", response)
        //     });
        // }
    }
}

Vue.createApp(EventHandling).mount('#app')