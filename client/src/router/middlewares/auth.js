import axios from "axios"

export default async function auth({ to, from, next, allMiddlewares, currentIndex }) {
    await axios.get("/api/user/user-check").then((response) => {
        if (response.status < 300 && response.status >= 200) {
            console.log("User check found")
            try {
                if (currentIndex >= allMiddlewares.length - 1 ) {
                    next()
                } else {
                    allMiddlewares[currentIndex+1]
                }
            } catch (error) {
                next()
            }  
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