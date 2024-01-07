import Navbar from './navbar'
import Footer from './footer'

function Layout({children}: {children: any}) {
  return(
    <div className="h-screen flex flex-col justify-between">
      <Navbar/>
      <div className='flex grow justify-center content'>
        <div className="flex w-[1000px] py-4">
          <main className="w-full">{children}</main>
        </div>
      </div>
      <div className="flex justify-center py-2">
        <div className="min-w-[1000px] text-sm">
          <Footer/>
        </div>
      </div>
    </div>
  )
}

export default Layout
