import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'

import Layout from '../components/layout'

function AuthPage() {
  const navigate = useNavigate()
  const [loginData, setLoginData] = useState({
    username: "", 
    password: ""
  })

  const [signupData, setSignupData]= useState({
    username: "",
    email: "",
    password: "",
    passwordMatch: ""
  })

  const handleLogin = async (e: React.MouseEvent) => {
    e.preventDefault()
    try {
      await axios.post(`${import.meta.env.VITE_API_URL}/api/public/login`, loginData, {
        withCredentials: true
      })
      navigate("/")
      return
    } catch(e) {
      console.error(e)
      alert("Credentials failed")
      return
    }
  }
  
  const handleSignup = async (e: React.MouseEvent) => {
    e.preventDefault()
    try {
      if(signupData.password != signupData.passwordMatch) {
        alert("Passwords do not match")
        throw new Error("Passwords do not match")
      }
      await axios.post(`${import.meta.env.VITE_API_URL}/api/public/signup`, signupData, {
        withCredentials: true
      })
      return
    } catch(e) {
      console.error(e)
    }

    // Should navigate to a page where it verifies that a confirmation email was sent
    navigate("/", { replace: true})
  }
  
  const inputClass = "border border-grey-300 px-2 py-2 rounded"
  const buttonClass = "bg-blue-500 rounded-2xl text-white py-1"

  return(
      <Layout>
        <div className="flex justify-center w-full">
          <div className="flex flex-col justify-center">
            <div className="bg-white rounded space-y-10 px-8 py-8 h-fit w-[350px]">
              <form className="flex flex-col space-y-3">
                <h1 className="text-2xl">Login</h1>
                <input 
                  type="text" 
                  placeholder="username"
                  name="username"
                  className={inputClass}
                  onChange={e => {
                    setLoginData({...loginData, username: e.target.value})
                  }}
                /> 
                <input 
                  type="password" 
                  placeholder="password"
                  name="password"
                  className={inputClass}
                  onChange={e => setLoginData({...loginData, password:e.target.value})}
                /> 
                <button
                  className={buttonClass}
                  onClick={handleLogin}
                >
                  Login</button>
              </form>
              <form className="flex flex-col space-y-3">
                <h1 className="text-2xl">Signup</h1>
                <input 
                  type="text" 
                  placeholder="email" 
                  className={inputClass}
                  onChange={e => setSignupData({...signupData, email: e.target.value})}
                /> 
                <input 
                  type="text" 
                  placeholder="username" 
                  className={inputClass}
                  onChange={e => setSignupData({...signupData, username: e.target.value})}
                /> 
                <input 
                  type="password" 
                  placeholder="password" 
                  className={inputClass}
                  onChange={e => setSignupData({...signupData, password: e.target.value})}
                /> 
                <input 
                  type="password" 
                  placeholder="match password" 
                  className={inputClass}
                  onChange={e => setSignupData({...signupData, passwordMatch: e.target.value})}
                /> 
                <button
                  className={buttonClass}
                  onClick={handleSignup}>
                  Signup
                </button>
              </form>
            </div>
          </div>
        </div>
      </Layout>
  )
}

export default AuthPage
