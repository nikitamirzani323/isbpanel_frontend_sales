<script>
    import { link } from "svelte-spa-router";
    async function handleLogout() {
        localStorage.clear();
        window.location.href = "/";
    }
  </script>
<div class="sticky top-0 z-30 flex h-16 w-full justify-center bg-opacity-90 backdrop-blur transition-all duration-100 
    bg-[#1a73e8] text-base-content shadow-sm">
    <div class="navbar bg-base-100">
        <div class="navbar-start">
          <div class="dropdown">
            <label tabindex="0" class="btn btn-ghost lg:hidden">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h8m-8 6h16" /></svg>
            </label>
            <ul tabindex="0" class="menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52">
              <li><a href="/">DASBOARD</a></li>
            
            </ul>
          </div>
          <a class="btn btn-ghost normal-case text-xl" href="/">SALES</a>
        </div>
        <div class="navbar-center hidden lg:flex">
          
        </div>
        <div class="navbar-end">
            <button
                on:click={() => {
                    handleLogout();
                }}
                class="btn btn-ghost btn-circle hover:bg-[#e6f7ff]">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-black fill-current" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                </svg>
            </button>
        </div>
      </div>
    
</div>