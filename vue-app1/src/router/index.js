import { createRouter, createWebHistory } from 'vue-router'
import Security from './../components/security'
import Body from './../components/Body.vue'
import Login from './../components/Login.vue'
import Books from './../components/Books.vue'
import Book from './../components/Book.vue'
import BooksAdmin from './../components/BooksAdmin.vue'
import BookEdit from './../components/BookEdit.vue'
import Users from './../components/Users.vue'
import User from './../components/UserEdit.vue'

const routes = [
    {
        // public
        path: '/',
        name: 'Home',
        component: Body,
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
    }, 
    {
        path: '/books',
        name: 'Books',
        component: Books,        
    },
    {
        // secured
        path: '/books/:bookName',
        name: 'Book',
        component: Book,   
        beforeEnter: Security.requireToken,        
    },
    {
        // secured
        path: '/admin/books/',
        name: 'Books admin',
        component: BooksAdmin,   
        beforeEnter: Security.requireToken,
    },
    {
        // secured
        path: '/admin/books/:bookId',
        name: 'BookEdit',
        component: BookEdit,  
        beforeEnter: Security.requireToken,      
    },
    {
        // secured
        path: '/admin/users',
        name: 'Users',
        component: Users,  
        beforeEnter: Security.requireToken,      
    },
    {
        // secured
        path: '/admin/users/:userId',
        name: 'User',
        component: User,  
        beforeEnter: Security.requireToken,      
    },
]

const router = createRouter({history: createWebHistory(), routes})
router.beforeEach(() => {
    Security.checkToken();
})


export default router