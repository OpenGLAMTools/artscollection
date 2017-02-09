const Collections = {
    template: `<a v-on:click="getData">load</a><input v-model="a">`,
    data: () => {
        this.getData();
        return {
            a: "asd"
        }
    },
    ready: () =>{
        this.getData();
    },
    props: ['$http'],
    methods: {
        getData: () => {
            Vue.http.get("/collection").then(
                (res) => {
                    console.log(res);
                }
            )
        }
    }
}


const Foo = { template: `<div>foo:<br>
<router-link to="/foo/foo">foo foo</router-link>
<router-view></router-view></div>` }
const Foofoo = { template: '<div>foooo</div>' }
const Bar = { template: '<div>bar</div>' }

// 2. Define some routes
// Each route should map to a component. The "component" can
// either be an actual component constructor created via
// Vue.extend(), or just a component options object.
// We'll talk about nested routes later.
const routes = [
  { path: '/foo', 
  component: Foo,
  children:[{ path: '/foo/foo', component: Foofoo }] },
  { path: '/bar', component: Bar },
   { path: '/collection', component: Collections }
]

// 3. Create the router instance and pass the `routes` option
// You can pass in additional options here, but let's
// keep it simple for now.
const router = new VueRouter({
  routes // short for routes: routes
})

// 4. Create and mount the root instance.
// Make sure to inject the router with the router option to make the
// whole app router-aware.
const app = new Vue({
  router
}).$mount('#app')