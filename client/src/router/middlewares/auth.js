import axios from "axios"

export default async function auth({ to, from, next }) {
    await axios.get("/api/user/user-check").then((response) => {
        if (response.status < 300 && response.status >= 200) {
            console.log("User check found")
            next()
        }
    }).catch((error) => {
        console.log(error.response)

        if (error.response.status !== 200) {
            console.log(from)
            if (from.name !== 'login' && to.name !== 'login') {
                console.log("REdirect stin login")

                next({ path: "/admin/login" })
            }
        }
    })


}