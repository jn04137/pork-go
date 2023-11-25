<script>
  let instance = new Date()
</script>

<div class="bg-white rounded px-4 py-2 shadow">
  <div class='py-2 text-sm'>
    <slot name="body">This is the body of the post.</slot>
  </div>
  <div class='flex flex-col text-xs'>
    <div>
      <slot name="author">Created by John Doe</slot>
    </div>
    <div class='italic'>
      <slot name="date">{instance.toLocaleString()}</slot>
    </div>
  </div>
</div>

<style>

</style>
