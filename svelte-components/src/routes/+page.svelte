<script>
  import ZooList from '../lib/List.comp.svelte'

  const animals = ['Cat', 'Dog', 'Fox', 'Pig', 'Horse', 'Chicken']
</script>

<main class="px-10">

  <header class="mb-10">
    <h1 class="text-5xl font-extrabold pt-5 pb-2">Welcome to your library project</h1>
    <section class="font-light">
      <p class="">Create your package using @sveltejs/package and preview/showcase your work with SvelteKit</p>
      <p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
    </section>
  </header>

  <section class="py-5">
    <div class="bg-gray-300 p-2 text-white border-b text-xl">Component: List</div>
    <div class="border">

      <ZooList title="Animals list" custom_items={animals} />

    </div>
  </section>

</main>
