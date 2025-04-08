import { createRouter, createWebHistory } from 'vue-router'
import Security from '../components/security'
import Body from '../components/Body.vue'
import Login from '../components/Login.vue'
import Books from '../components/Books.vue'
import Book from '../components/Book.vue'
import BooksAdmin from '../components/BooksAdmin.vue'
import BookEdit from '../components/BookEdit.vue'
import Users from '../components/Users.vue'
import User from '../components/UserEdit.vue'

const secureRoute = (path, name, component) => ({
    path, 
    name, 
    component, 
    beforeEnter: Security.requireToken,
});

const routes = [
    {path: '/', name: 'Home', component: Body },
    {path: '/login', name: 'Login', component: Login }, 
    {path: '/books',name: 'Books', component: Books },
    {path: '/books/:bookName',name: 'Book', component: Book },
    
    //secureRoute('/books/:bookName', 'Book', Book),
    secureRoute('/admin/books/', 'Manage Books', BooksAdmin),
    secureRoute('/admin/books/:bookId', 'BookEdit', BookEdit),
    secureRoute('/admin/users', 'Users', Users),
    secureRoute('/admin/users/:userId', 'User', User),
]

const router = createRouter({history: createWebHistory(), routes})

router.beforeEach(() => {
    Security.checkToken();
})

export default router