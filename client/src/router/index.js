import Vue from "vue";
import VueRouter from "vue-router";
import Dashboard from 'Views/pages/Dashboard';
import Editor from 'Views/pages/Editor';
import Login from 'Views/auth/Login';
import auth from "./middlewares/auth";

const routes = [
    {
        component: Login,
        path: '/admin/login',
        name: 'login',
        meta: {
            middlewares: null
        }

    },
    {
        component: Dashboard,
        path: '/admin',
        name: 'dashboard',
        meta: {
            middlewares: [auth]
        }
    },
    {
        component: Editor,
        path: '/admin/editor',
        name: 'editor',
        meta: {
            middlewares: [auth]
        }
    }
]

Vue.use(VueRouter);

const router = new VueRouter({
    // 4. Provide the history implementation to use. We are using the hash history for simplicity here.
    mode: "history",
    routes, // short for `routes: routes`
})

router.beforeEach((to, from, next) => {
    if (to.name === null){console.log("Redirect from guard"); next({ name: "dashboard" })}
    const middlewares = to.meta.middlewares
    const args = { to, from, next }
    // return next()
    if (!middlewares) {
        return next();
    }

    middlewares[0]({
        ...args, allMiddlewares: middlewares, currentIndex: 0
    })
})

export default router;