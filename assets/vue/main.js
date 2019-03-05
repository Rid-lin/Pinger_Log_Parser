const app = new Vue({
    el: "#app",
    data:{
        editAdress: null,
        adresses: {
            type: Array,
            default: function(){
                return []
            },
        },            
        adress: {
            IP: '',
            Note: '',
            SiteID: '',
            StatusNow:{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },
            StatusOfHour: [
            {
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },{
                Code: '',
                Background: '',
                NumPass:'',
                NumFail:'',
            },
            ],
        
        },
    },
    methods:{
        httpadress(value){
            return `http://${value}`;
        },
        deleteAdress(IP, i){
            fetch("http://127.0.0.1:8080/api/v1/adress"+"/"+IP,{
                method: "DELETE"
            })
            .then(() => {
                delete this.adresses[i]
                // this.adresses.splice(i, 1);
            })
        },
        updateAdress(adress){
            fetch("http://127.0.0.1:8080/api/v1/adress"+"/"+adress.IP,{
                body: JSON.stringify(adress),
                method: "PUT",
                headers : {
                    "Content-Type": "application/json",
                },
            })
            .then(() => {
                this.editAdress = null;
            })
        },
        saveConfig(){
            fetch("http://127.0.0.1:8080/api/v1/saveconf")
        },
        createAdress(){
            adress = {
                IP: this.IP,
                Note: this.Note,
                SiteID: this.SiteID,
                StatusNow:{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail
                },
                StatusOfHour: [
                {
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },{
                    Code: this.Code,
                    Background: this.Background,
                    NumPass:this.NumPass,
                    NumFail:this.NumFail,
                },
                ],
            
            }
    
            fetch("http://127.0.0.1:8080/api/v1/adress",{
                body: JSON.stringify(adress),
                method: "POST",
                headers : {
                    "Content-Type": "application/json",
                },
            })
            .then(() => {
                this.editAdress = null;
            })

        },
    },
    mounted(){
        fetch("http://127.0.0.1:8080/api/v1/adresses")
            .then(response => response.json())
            .then((data) => {
                this.adresses = data;
            })
    },
})