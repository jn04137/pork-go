import { Link } from 'react-router-dom'
import axios from 'axios'
import { useQuery } from '@tanstack/react-query'

export let navs = [
  { id: 0, endpoint: '/', page: 'Home' },
  //{ endpoint: '/about', page: 'About' },
  //{ id: 1, endpoint: '/contact', page: 'Contact' }
]

const isLoggedIn = async () => {
  try {
    const res = await axios.get(`${import.meta.env.VITE_API_URL}/api/public/isloggedin`, {
      withCredentials: true
    })
    return res.data
  } catch(err) {
    console.error(err)
  }
}

function Navbar() {
  const loggedInQuery = useQuery({
      queryKey: ['userLoggedIn'],
      queryFn: isLoggedIn,
      refetchOnWindowFocus: false,
      refetchInterval: 300000
  })


  return(
    <div className='flex justify-center shadow-lg bg-[#333333]'>
      <div className='flex justify-between py-2 w-[1000px] items-center'>
        <div className='text-lg font-bold'><Link to="/">satch3l</Link></div>
        <div className="flex space-x-4 items-center text-sm">
          <nav className='space-x-4'>
            {navs.map(nav => {
              return <Link to={nav.endpoint} key={nav.id}>{nav.page}</Link>
            })}
          </nav>
          {!loggedInQuery.isLoading && <AuthButton loggedIn={loggedInQuery.data.isLoggedIn}/>}
        </div>
      </div>
    </div>
  )
}

function AuthButton({loggedIn}: {loggedIn: boolean}) {
  if(loggedIn) {
    return <Link to="/myprofile">My Profile</Link>
  } else {
    return(
      <Link to="/authpage">
        <button className="bg-blue-500 text-white px-2 py-0.5 rounded-2xl shadow">Login / Signup</button>
      </Link>
    ) 
  }
}

export default Navbar
