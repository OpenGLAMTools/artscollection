const RenderString = Vue.component('render-string',{
    template: `<div class="field">
        <label>{{ field.Name }}</label> 
        <input v-model="storage[field.Key]"></div>`,
    created: function (){
        if (this.storage[this.field.Key] == null ){
            this.$set(this.storage, this.field.Key, "")
            }
    },
    props: {
        storage: Object,
        field: Object
    }
})
const RenderInteger = Vue.component('render-integer',{
    template: `<div class="field">
    <label>{{ field.Name }}</label> 
    <input type="number" v-model="storage[field.Key]"></div>`,
    created: function (){
        if (this.storage[this.field.Key] == null ){
            this.$set(this.storage, this.field.Key, 0)
            }
    },
    props: {
        storage: Object,
        field: Object
    }
})
const RenderBool = Vue.component('render-bool',{
    template: `<div class="field">
    <label>{{ field.Name }}</label>
     <input type="checkbox" v-model="storage[field.Key]"></div>`,
    created: function (){
        if (this.storage[this.field.Key] == null ){
            this.$set(this.storage, this.field.Key, false)
            }
    },
    props: {
        storage: Object,
        field: Object
    }
})
const RenderList = Vue.component('render-bool',{
    template: `<div>{{ field.Name }}: 
    <ul>
    <li v-for="(value,key) in storage[field.Key]"><input v-model="storage[field.Key][key]"></li>
    <li @click="addTag">Add Tag</li>
    </ul>
    </div>`,
    props: {
        storage: Object,
        field: Object
    },
    created: function (){
        if (this.storage[this.field.Key] == null ){
            this.$set(this.storage, this.field.Key, [])
            }
    },
    methods: {
        addTag: function (){
            if (this.storage[this.field.Key] == null ){
                this.storage[this.field.Key] = [];
            }
            this.storage[this.field.Key].push("");
        }
    }
})
const RenderField = Vue.component('render-field',{
    template: `<div>
    <component :is=field.Type :storage=storage :field=field></component></div>`,
    computed: {
        fieldType: function(){
            return this.field.Type
        }
    },
    components: {
        string: RenderString,
        int: RenderInteger,
        bool: RenderBool,
        list: RenderList
    },
    
    props: {
        field: Object,
        storage: Object
    }
})
const Home = Vue.component('home', {
    data: function () {
        return {
            collections: {}
        }
    },
    template: `
    <div class="ui grid">
        <div class="five wide column">
            <div class="ui vertical pointing menu">
            <span  v-for="(col, key) in collections">
                <router-link class="item" :to="{ name: 'collection', params: { cid: key}}">
                {{key}}
                </router-link>
            </span>
            </div>
             
        </div>
        <div class="ten wide column">
            <router-view></router-view>
        </div>  
     
    </div>
    `,
    mounted: function () {
        this.$http.get('/collection').then(
            (res) => {
                this.collections = res.body;
            }
        );
        //this.collections = data;
    }
})

const Collection = Vue.component('collection', {
    template: `
    <div class="ui grid">
        <div class="six wide column">
            <h2>Collection</h2>
            <ul>
            <li v-for="(item, key) in items.Storages">
            <router-link :to="{ name: 'item', params: { iid: key}}">
                {{key}}
                </router-link>
            </li>
            </ul>
       </div>
       <div class="six wide column">
            <router-view></router-view>
        </div>
    </div>
    `,
    data: function () {
        return {
            items: { Storages: null }
        }
    },
    methods: {
        fetchData: function () {
            this.$http.get('/collection/' + this.cid).then(
                (res) => {
                    this.items = res.body;
                }
            );
        }
    },
    created: function () {
        this.fetchData()
    },
    watch: {
        '$route': 'fetchData'
    },
    props: ['cid']
})

const Item = Vue.component('item', {
    template: `
    <div>
    <form class="ui form">
        <h3>{{ iid }}</h3>
        
        <div v-for="f in item.fields">
        <render-field :field=f :storage=item[f.Type] ></render-field>
        </div>
        
        <button class="ui button" @click="saveData">Save Data</button>
    </form>
    </div>`,
    data: function () {
        return {
            item: {fields:[]}
        }
    },
    methods: {
        fetchData: function () {
            this.$http.get('/collection/' + this.cid + '/' + this.iid).then(
                (res) => {
                    this.item = res.body;
                }
            );
        },
        saveData: function () {
            this.$http.post('/collection/' + this.cid + '/' + this.iid,this.item);
        }
    },
    created: function () {
        this.fetchData()
    },
    watch: {
        '$route': 'fetchData'
    },
    props: ['cid', 'iid']
})

const routes = [
    {
        path: '/',
        name: 'home',
        component: Home,
        children: [
            {
                path: ':cid',
                name: 'collection',
                props: true,
                component: Collection,
                children: [
                    {
                        path: ':iid',
                        name: 'item',
                        props: true,
                        component: Item
                    }
                ]
            }]
    },
]


const router = new VueRouter({
    routes: routes
})


const app = new Vue({
    router
}).$mount('#app')