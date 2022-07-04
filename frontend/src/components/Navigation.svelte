<script>
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import { link } from "svelte-spa-router";
    import Loader from '../components/Loader.svelte' 
    import Input_custom from '../components/Input.svelte' 
    import Modal_popup from '../components/Modal_popup.svelte' 
    import Modal_alert from '../components/Modal_alert.svelte' 
    export let path_api = ""
    let master = localStorage.getItem("master");
    let token = localStorage.getItem("token");
    let isModal_Form_password = false
    let isModalLoading = false
    let isModalNotif = false
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let msg_error = "";
    const schema = yup.object().shape({
        password: yup.string().required().min(4,"Password must be at least 4 characters").max(50,"Password must be at most 50 characters")
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            password: "",
        },
        validationSchema: schema,
        onSubmit:(values) => {
            handleSave(values.password)
        }
    })
    async function handleSave(password) {
        msg_error = "";
        loader_class = "inline-block"
        loader_msg = "Sending..."
        const res = await fetch(path_api+"api/changepassword", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                password: password,
            }),
        });
        const json = await res.json();
        if(!res.ok){
            loader_msg = "System Mengalami Trouble"
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
        }else{
            if (json.status == 200) {
                loader_msg = json.message
            } else if (json.status == 403) {
                loader_msg = json.message
            } else {
                loader_msg = json.message;
            }
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
        }
        isModalLoading = false;
        isModal_Form_password = false;
    }
    const handlePassword = () => {
      $form.password = "";
      isModal_Form_password = true
    };
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
          <span class="pr-5 font-semibold select-none">
            Welcome, {master}
          </span>
          <button
            on:click={() => {
                handlePassword();
            }}
            class="btn btn-ghost btn-circle hover:bg-[#e6f7ff]">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z" />
            </svg>
          </button>
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
<Loader loader_class="{loader_class}" loader_msg="{loader_msg}" />
<input type="checkbox" id="my-modal-formnew" class="modal-toggle" bind:checked={isModal_Form_password}>
<Modal_popup
    modal_popup_id="my-modal-formnew"
    modal_popup_title="Change Password"
    modal_popup_class="select-none max-w-full lg:max-w-xl overflow-hidden">
    <slot:template slot="modalpopup_body">
        <div class="flex flex-col w-full gap-2">
          <div class="relative form-control mt-2">
              <Input_custom
                  input_onchange="{handleChange}"
                  input_autofocus={false}
                  input_required={true}
                  input_tipe="password"
                  input_attr="password"
                  input_invalid={$errors.password.length > 0}
                  bind:value={$form.password}
                  input_id="password"
                  input_placeholder="Password"
                  />
              {#if $errors.password}
                  <small class="text-pink-600 text-[11px]">{$errors.password}</small>
              {/if}
          </div>
            <button on:click={() => {
                handleSubmit();
            }} class="btn btn-primary">Update</button>
        </div>
    </slot:template>
</Modal_popup>

<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />