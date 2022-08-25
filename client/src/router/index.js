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
    mode: "history",
    routes,
})

router.beforeEach((to, from, next) => {
    if (to.name === null) { next({ name: "dashboard" }) }
    const middlewares = to.meta.middlewares
    const args = { to, from, next }

    if (!middlewares) {
        return next();
    }

    middlewares[0]({
        ...args, allMiddlewares: middlewares, currentIndex: 0
    })
})

export default router;