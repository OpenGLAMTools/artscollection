const Home = Vue.component('home', {
    data: function () {
        return {
            collections: {}
        }
    },
    template: `
    <div>
    <h1>Home</h1>
    <ul>
        <li v-for="(col, key) in collections">
        <router-link :to="{ name: 'collection', params: { cid: key}}">
        {{key}}
        </router-link>
        </li>
    </ul>
    <router-view></router-view>
    </div>`,
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
    template: `<div>
    <h2>Collection</h2>
    <ul>
    <li v-for="(item, key) in items.Storages">
    <router-link :to="{ name: 'item', params: { iid: key}}">
        {{key}}
        </router-link>
    </li>
    </ul>
    <router-view></router-view></div>`,
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
    template: `<div><h3>Item</h3>
    <h4>{{ iid }}</h4>
    <ul>
    <li v-for="f in item.fields">{{f.Name}}</li>
    </ul>
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