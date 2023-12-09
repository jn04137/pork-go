<script>
  import './app.css'
  import { Router, Route } from 'svelte-routing'
  import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query'

  import Home from './pages/home.svelte'
  import About from './pages/about.svelte'
  import Contact from './pages/contact.svelte'
  import MyProfile from './pages/my_profile.svelte'
  import CreatePost from './pages/create_post.svelte'
  import LoginPage from './pages/login.svelte'
  import SignupPage from './pages/signup.svelte'
  import ViewPostPage from './pages/viewpost.svelte'
  import UserDashboard from './pages/user_dashboard.svelte'
  import Admin from './pages/admin.svelte'

  export let url = ""
  const queryClient = new QueryClient()
</script>

<QueryClientProvider client={queryClient}>
  <main>
    <Router {url}>
      <div>
        <Route path="/" component={Home}/>
        <Route path="/about" component={About}/>
        <Route path="/contact" component={Contact}/>
        <Route path="/myprofile" component={MyProfile}/>
        <Route path="/createpost" component={CreatePost}/>
        <Route path="/login" component={LoginPage}/>
        <Route path="/signup" component={SignupPage}/>
        <Route path="/viewpost/:id" let:params>
          <ViewPostPage postId={params.id}/>
        </Route>
        <Route path="/userdashboard" component={UserDashboard}/>
        <Route path="/admin" component={Admin}/>
      </div>
    </Router>
  </main>
</QueryClientProvider>

<style>
</style>
