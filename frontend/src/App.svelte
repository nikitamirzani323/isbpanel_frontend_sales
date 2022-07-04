<script>
	import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
	import Navigation from '../src/components/Navigation.svelte'
	import NotFound from '../src/pages/NotFound.svelte'
	import Login from '../src/pages/Login.svelte'
	import Admin from '../src/pages/admin/Admin.svelte'
	import Tailwindcss from './Tailwindcss.svelte'
	export let path_api = "";
	export let font_size = "";
  	let token = localStorage.getItem("token");
	let routes = "";
	let isNav = false;
  	if (token == "" || token == null) {
    // isNav = true;
		routes = {
			"/": wrap({
        		props: {
					path_api: path_api,
				},
				component: Login,
			}),
			"*": NotFound,
		};
	} else {
		isNav = true;
			routes = {
				"/": wrap({
					props: {
						path_api: path_api,
						font_size: font_size,
					},
					component: Admin,
				}),
		
			
			"*": NotFound,
			};
  }
</script>
{#if isNav}
<main class="flex flex-col  h-screen w-full bg-[#f0f2f5]">
	<Navigation path_api="{path_api}" />
	<div class="justify-center ">
    	<div class="w-full  mt-5">
      		<Router {routes} />
    	</div>
  	</div>
</main>
{:else}
<div class="container mx-auto">
  <Router {routes} />
</div>
{/if}
<Tailwindcss />