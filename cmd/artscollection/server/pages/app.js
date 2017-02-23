const RenderString = Vue.component('render-string', {
    template: `<div class="field">
        <label>{{ field.Name }}</label> 
        <input v-model="storage[field.Key]"></div>`,
    created: function () {
        if (this.storage[this.field.Key] == null) {
            this.$set(this.storage, this.field.Key, "")
        }
    },
    props: {
        storage: Object,
        field: Object
    }
})
const RenderInteger = Vue.component('render-integer', {
    template: `<div class="field">
    <label>{{ field.Name }}</label> 
    <input type="number" v-model="storage[field.Key]"></div>`,
    created: function () {
        if (this.storage[this.field.Key] == null) {
            this.$set(this.storage, this.field.Key, 0)
        }
    },
    props: {
        storage: Object,
        field: Object
    }
})
const RenderBool = Vue.component('render-bool', {
    template: `<div class="field">
    <label>{{ field.Name }}</label>
     <input type="checkbox" v-model="storage[field.Key]"></div>`,
    created: function () {
        if (this.storage[this.field.Key] == null) {
            this.$set(this.storage, this.field.Key, false)
        }
    },
    props: {
        storage: Object,
        field: Object
    }
})
const RenderList = Vue.component('render-list', {
    template: `<div>{{ field.Name }}: 
    <div class="ui label" v-for="(value,key) in storage[field.Key]">
    {{ storage[field.Key][key] }} <i class="delete icon" @click="removeTag(key)"></i>
    </div><br>
    <div class="ui right labeled input">
  <input type="text" v-model="newTagVal">
  <a class="ui tag label" @click="addTag()">
    Add {{ field.Name }}
  </a>
</div>
    </div>`,
    data: function (){
        return {
            newTagVal: ""
        }
    },
    props: {
        storage: Object,
        field: Object
    },
    created: function () {
        if (this.storage[this.field.Key] == null) {
            this.$set(this.storage, this.field.Key, [])
        }
    },
    methods: {
        addTag: function () {
            if (this.storage[this.field.Key] == null) {
                this.storage[this.field.Key] = [];
            }
            this.storage[this.field.Key].push(this.newTagVal);
            this.newTagVal = "";
        },
        removeTag: function (key) {
            this.storage[this.field.Key].splice(key,1);
        }
    }
})
const RenderField = Vue.component('render-field', {
    template: `<div>
    <component :is=field.Type :storage=storage :field=field></component></div>`,
    computed: {
        fieldType: function () {
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
const ListCollections = Vue.component('listcollections', {
    template: `
    <div class="menu">
        <router-link 
                    v-for="(col, key) in collections"
                    active-class="active"
                    class="item" 
                    :to="{ name: 'collection', params: { cid: key}}">
                {{key}}
                </router-link>
    </div>
    `,
    data: function () {
        return {
            collections: {}
        }
    },
    mounted: function () {
        this.$http.get('/collection').then(
            (res) => {
                this.collections = res.body;
            }
        );
    }
});
const Home = Vue.component('home', {
    data: function () {
        return {
            collections: {}
        }
    },
    template: `
    <div class="ui stackable grid">
        <div class="two wide column">
            <div class="ui relaxed divided selection list">
            <div
                class="item"  
                v-for="(col, key) in collections">
                <router-link 
                    active-class="active"
                    class="header" 
                    :to="{ name: 'collection', params: { cid: key}}">
                {{key}}
                </router-link>
            </div>
            </div>
             
        </div>
        <div class="ten wide column">
            <router-view></router-view>
        </div>  
     
    </div>
    `,
    props: ['cid'],
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
    <div class="ui stackable grid">
        <div class="three wide column">
            <div class="ui relaxed divided selection list">
            <div
                class="item"
                v-bind:class="{ active: key==iid }" 
                v-for="(item, key) in items.Storages">
                <router-link 
                    active-class="active"
                    class="header" 
                    :to="{ name: 'item', params: { iid: key}}">
                    {{key}}
                </router-link>
            </div>
            </div>

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
    props: ['cid','iid']
})

const Item = Vue.component('item', {
    template: `
    <div>
    <form class="ui small form">
        <h3>{{ iid }}</h3>
        <div v-for="f in item.fields">
        <render-field :field=f :storage=item[f.Type] ></render-field>
        </div>
        <button class="ui button" @click="saveData">Save Data</button>
    </form>
    </div>`,
    data: function () {
        return {
            item: { fields: [] }
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
            this.$http.post('/collection/' + this.cid + '/' + this.iid, this.item);
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

const Breadcrumb = Vue.component('breadcrumb', {
    template: `
<div class="ui breadcrumb">
    <router-link class="section" to="/">Home</router-link>
    <span class="divider" v-if="cid">/</span>
    <router-link
                v-if="cid" 
                class="section" 
                :to="{ name: 'collection', params: { cid: cid}}">
                {{cid}}
                </router-link>
    <span class="divider" v-if="iid">/</span>
    <div class="active section" v-if="iid">{{ iid }}</div>
</div>
    `,
    props: ['cid', 'iid']
})

const routes = [
    {
        path: '/',
        name: 'home',
        components: {
            default: Home,
            breadcrumb: Breadcrumb,
            listcollections: ListCollections
        },
        children: [
        ]
    },
    {
        path: '/:cid',
        name: 'collection',
        components: {
            default: Collection,
            breadcrumb: Breadcrumb,
            listcollections: ListCollections
        },
        props: {
            default: true,
            breadcrumb: true
        },
        children: [
            {
                path: '/:cid/:iid',
                name: 'item',
                component: Item,
                props: true
            }
        ]
    }
]


const router = new VueRouter({
    routes: routes
})


const app = new Vue({
    router,
    data: {
        collectionName: "",
        itemName: ""
    }
}).$mount('#app')


$(document)
    .ready(function () {
        $('.ui.dropdown')
            .dropdown()
            ;
    })
    ;