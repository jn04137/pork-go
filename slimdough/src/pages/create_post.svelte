<script>
  import Layout from '../lib/components/layout.svelte'
  import TipTap from '../lib/components/editor.svelte'
  import { navigate } from 'svelte-routing'
  import axios from 'axios'

  let editor

  let postData = {
    title: "",
    body: null
  }

  const handlePublish = async () => {
    if(postData.title.length === 0) {
      alert("Please enter a title")
      throw new Error("Title wasn't supplied")
    };
    try {
      await axios.post("http://localhost:8000/api/createpost", postData, {
        withCredentials: true
      })
      console.log(postData)
    } catch(e) {
      console.error(e)
    }
    navigate("/", { replace: true })
  }

</script>

<Layout>
  <body slot="content" class="w-full">
    <div class="flex justify-center pt-10">
      <form class="bg-white rounded p-4 w-[650px] space-y-2 shadow-lg">
        <input 
          placeholder="title" 
          class="text-xl rounded py-0.5 px-2 border border-grey-300 w-full"
          bind:value={postData.title}
          required
        />
        <TipTap
          bind:editorHTML={postData.body}
        />
        <div class="flex justify-end">
          <button
            on:click|preventDefault={() => handlePublish()}
            class="bg-blue-500 text-white px-4 py-1 rounded-2xl shadow">
            Publish
          </button>
        </div>
      </form>
    </div>
  </body>
</Layout>

