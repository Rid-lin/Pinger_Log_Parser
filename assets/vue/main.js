const app = new Vue({
    el: "#app",
    data:{
        editAdress: null,
        adresses: [],            
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
    template:`
    <div> 
        <div v-if="editAdress === 100">
            IP-адрес: <input v-on:keyup.13="createAdress()" v-model="IP"    />
            <br>Описание: <input v-on:keyup.13="createAdress()" v-model="Note" />
            <br>SiteID: <input v-on:keyup.13="createAdress()" v-model="SiteID" />
            <button v-on:click="createAdress()">Сохранить</button>
            <button v-on:click="editAdress = null">Отмена</button>
        </div>
        <div v-else>
            <button v-on:click="editAdress = 100">Добавить адрес</button>
            <button v-on:click="">Сохранить конфиг</button>
        </div>
        <table>
            <thead>
                <tr v-for="adress, i  in adresses" v-if="adress.IP == 'IP адрес'">
                    <th>{{adress.IP}}</th>
                    <th>Редак.</th>
                    <th>{{adress.Note}}</th>
                    <th>{{adress.SiteID}}</th>
                    <th>{{adress.StatusNow.Code}}</th><th v-for="item in adress.StatusOfHour">{{item.Code}}</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="adress, i  in adresses" v-if="adress.IP != 'IP адрес'">
                    <td v-if="editAdress === adress.IP">
                        <input v-on:keyup.13="updateAdress(adress)" v-model="adress.IP" />
                    </td>
                    <td v-else><a :href="httpadress(adress.IP)">{{adress.IP}}</a></td>

                    <td v-if="editAdress === adress.IP">
                        <button v-on:click="updateAdress(adress)">Сохранить</button>
                        <button v-on:click="editAdress = null">Отмена</button>
                        <button v-on:click="deleteAdress(adress.IP, i)">Удалить</button>
                    </td>
                    <td v-else>
                        <button v-on:click="editAdress = adress.IP">Ред.</button>
                    </td>

                    <td v-if="editAdress === adress.IP">
                        <input v-on:keyup.13="updateAdress(adress)" v-model="adress.Note" />
                    </td>
                    <td  class="note" v-else>{{adress.Note}}</td>

                    <td v-if="editAdress === adress.IP">
                        <input v-on:keyup.13="updateAdress(adress)" v-model="adress.SiteID" />
                    </td>
                    <td  class="small-font" v-else>{{adress.SiteID}}</td>

                    <td v-bind:class="adress.StatusNow.Background">{{adress.StatusNow.Code}}</td><td v-for="item in adress.StatusOfHour"  v-bind:class="item.Background">{{item.Code}}</td>
                </tr>
            </tbody>
        </table>
    </div>
    `
})