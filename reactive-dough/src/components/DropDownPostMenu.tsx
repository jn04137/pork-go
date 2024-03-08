import { 
	useState,
	Fragment,
	useRef
} from 'react'
import { useNavigate } from 'react-router-dom'

import { PiDotsThreeOutlineFill } from 'react-icons/pi'

import { 
	Transition, 
	Menu,
	Dialog
} from '@headlessui/react'

import axios from 'axios'

const apiURL = import.meta.env.VITE_API_URL

function DropDownPostMenu({postId}: {
  postId: number
}) {
  const [open, setOpen] = useState(false)
	const navigate = useNavigate()

  // Send a request to the backend to delete the post
  async function handleDeleteClick(e: React.SyntheticEvent) {
    e.preventDefault()
    try {
      setOpen(true)
    } catch(e) {
      console.error(e)
    }
  }

	async function handleEditClick(e: React.SyntheticEvent) {
		e.preventDefault()
		try {
			navigate(`/editpost/${postId}`, { replace: true })
		} catch(e) {
			console.error(e)
		}
	}

  return(
    <Menu as="div" className="relative inline-block text-left text-sm">
      <div>
        <Menu.Button
          className="inline-flex w-full justify-center gap-x-1.5 rounded-md px-3 py-1
          text-sm font-semibold ring-1 ring-inset ring-gray-300 
          dark:hover:bg-[#121212]"
        >
          <PiDotsThreeOutlineFill/>
        </Menu.Button>
      </div>
      <Transition
        as={Fragment}
      >
        <Menu.Items
          className={`absolute right-0 z-10 mt-2 w-24 origin-top-right rounded-md
          ring-1 ring-black ring-opacity-5 focus:outline-none bg-[#121212] px-2 py-2 space-y-2`}
        >
          <Menu.Item>
            <a 
              className="block hover:cursor-pointer"
							onClick={handleDeleteClick}
            >Delete Post</a>
          </Menu.Item>
          <Menu.Item>
            <a 
              className="block hover:cursor-pointer"
							onClick={handleEditClick}
            >Edit Post</a>
          </Menu.Item>
        </Menu.Items> 
      </Transition>
      <ConfirmDeletePostDialog open={open} setOpen={setOpen} postId={postId}/>
    </Menu>
  )
}

function ConfirmDeletePostDialog({open, setOpen, postId}: {
    open: boolean,
    setOpen: any,
    postId: number
}) {
	const cancelButtonRef = useRef(null)
  const navigate = useNavigate()

	async function handleDeleteClick(e: React.MouseEvent<HTMLElement>) {
		e.preventDefault()

    try {
			console.log("This is the postId:", postId)
      const response = await axios.delete(`${apiURL}/api/protected/post/delete`, {
				withCredentials: true,
				data: {
					postId: postId
				}
      }) 
      if(response.status === 200) navigate("/", { replace: true })
    } catch(e) {
      console.log(e)
    }
  }

  return(
    <Transition.Root show={open} as={Fragment}>
      <Dialog as="div" className="relative z-10" initialFocus={cancelButtonRef} onClose={setOpen}>
        <Transition.Child
          as={Fragment}
          enter="ease-out duration-300"
          enterFrom="opacity-0"
          enterTo="opacity-100"
          leave="ease-in duration-200"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
        </Transition.Child>

        <div className="fixed inset-0 z-10 w-screen overflow-y-auto">
          <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
            <Transition.Child
              as={Fragment}
              enter="ease-out duration-300"
              enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
              enterTo="opacity-100 translate-y-0 sm:scale-100"
              leave="ease-in duration-200"
              leaveFrom="opacity-100 translate-y-0 sm:scale-100"
              leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
            >
              <Dialog.Panel className="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
                <div className="bg-[#212121] px-4 pb-4 pt-5 sm:p-6 sm:pb-4">
                  <div className="sm:flex sm:items-start">
                    <div className="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
                      <Dialog.Title as="h3" className="text-base font-semibold leading-6 text-white">
                        Delete Post
                      </Dialog.Title>
                      <div className="mt-2">
                        <p className="text-sm text-[#AAAAAA]">
                          Are you sure you want to delete this post?<br/>This will be a permanent action.
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
                <div className="bg-[#3f3f3f] px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
                  <button
                    type="button"
                    className="inline-flex w-full justify-center rounded-md bg-red-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 sm:ml-3 sm:w-auto"
                    onClick={handleDeleteClick}
                  >
                    Delete
                  </button>
                  <button
                    type="button"
                    className="mt-3 inline-flex w-full justify-center rounded-md bg-[#5c5c5c] px-3 py-2 text-sm font-semibold text-white hover:bg-[#212121] sm:mt-0 sm:w-auto"
                    onClick={() => setOpen(false)}
                    ref={cancelButtonRef}
                  >
                    Cancel
                  </button>
                </div>
              </Dialog.Panel>
            </Transition.Child>
          </div>
        </div>
      </Dialog>
    </Transition.Root>
  )
}

export default DropDownPostMenu
