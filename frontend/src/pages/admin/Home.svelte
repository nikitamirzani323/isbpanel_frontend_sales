<script>
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import Input_custom from '../../components/Input.svelte' 
    import Modal_alert from '../../components/Modal_alert.svelte' 
    import Modal_popup from '../../components/Modal_popup.svelte' 
    import Loader from '../../components/Loader.svelte' 
    import Panel from '../../components/Panel_default.svelte' 
    import Panel_info from '../../components/Panel_info.svelte'

    export let path_api = "";
    export let font_size = "";
    export let token = "";
    export let listHome = [];
    export let admin_listrule = [];
    export let totalrecord = 0;

    let page = "CRM PROCESS";
    let sData = "New";
    let isModal_Form_New = false
    let isModal_Form_Listipaddress = false
    let isModalLoading = false
    let isModalNotif = false
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_class = "btn btn-primary"
    let msg_error = "";

    let status_crm_satu = "";
    let status_crm_dua = "";
    let img_deposit = "deposit.png"
    let img_reject = "reject.svg"
    let img_noanswer = "no-phone.svg"
    let img_check = "check.svg"
    let panel_deposit = false
    let panel_reject = false
    let panel_noanswer = false
    let website_field = ""
    let deposit_field = 0

    let admin_listip = [];
    let tab_menu_1 = "bg-sky-600 text-white"
    let tab_menu_2 = ""
    let panel_edit = true
    let panel_iplist = false
    let admin_tipe = "ADMIN";
    let searchHome = "";
    let filterHome = [];
    let form_field_ipaddress = "";
    let isInput_username_enabled = true;
    let admin_create_field = ""
    let admin_update_field = ""
    let dispatch = createEventDispatcher();
    const schema = yup.object().shape({
        admin_username_field: yup
            .string()
            .required("Username is Required")
            .matches(
                /^[a-zA-z0-9]+$/,
                "Username must Character A-Z or a-z or 1-9"
            )
            .min(4, "Username must be at least 4 Character")
            .max(20, "Username must be at most 20 Character"),
        admin_password_field: yup.string(),
        admin_name_field: yup
            .string()
            .required("Nama is Required")
            .matches(
                /^[a-zA-z0-9]+$/,
                "Nama must Character A-Z or a-z or 1-9"
            )
            .min(4, "Nama must be at least 4 Character")
            .max(20, "Nama must be at most 20 Character"),
        admin_idrule_field: yup.number().required("Admin Rule is Required"),
        admin_status_field: yup.string().required("Status is Required"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            admin_username_field: "",
            admin_password_field: "",
            admin_name_field: "",
            admin_idrule_field: "0",
            admin_status_field: "",
        },
        validationSchema: schema,
        onSubmit: (values) => {
            SaveTransaksi(
                values.admin_username_field,
                values.admin_password_field,
                values.admin_name_field,
                values.admin_idrule_field,
                values.admin_status_field
            );
        },
    });
    async function SaveTransaksi(username, password, name, rule,status) {
        let flag = true;
        msg_error = "";
        const regexExp = /^[a-zA-z0-9]+$/gi;
        let flag_password = regexExp.test(password)
        if(password != ""){
            if(!flag_password){
                flag = false;
                msg_error += "The Format Password invalid\n Password must Character A-Z or a-z or 1-9";
            }
        }
        if (rule == "0") {
            flag = false;
            msg_error += "The Admin Rule is required";
        }
        if(status == ""){
            flag = false;
            msg_error += "The Status is required";
        }
        if (flag) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/saveadmin", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    idruleadmin: parseInt(rule),
                    page: "ADMIN-SAVE",
                    username: username,
                    password: password,
                    nama: name,
                    status: status,
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
                    if(sData == "New"){
                        $form.admin_username_field = "";
                        $form.admin_password_field = "";
                        $form.admin_name_field = "";
                        $form.admin_idrule_field = "0";
                        $form.admin_status_field = "";
                    }
                } else if (json.status == 403) {
                    loader_msg = json.message
                } else {
                    loader_msg = json.message;
                }
                buttonLoading_class = "btn btn-primary"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
                RefreshHalaman();
            }
        } else {
            alert(msg_error);
        }
    }
    async function SaveIpaddress() {
        let flag = true;
        let totaliplist = admin_listip.length;
        msg_error = "";
        if (form_field_ipaddress == "") {
            flag = false;
            msg_error += "The Ipaddress is required\n";
        }
        const regexExp = /^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/gi;
        let flag_ip = regexExp.test(form_field_ipaddress)
        if(!flag_ip){
            flag = false;
            msg_error += "The Format Ipaddress invalid\n";
        }
        if(totaliplist > 5){
            flag = false;
            msg_error += "Maximal 5 Ipaddress Active\n";
        }
        if (flag) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/saveadminiplist", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sData: "New",
                    page: "ADMIN-SAVE",
                    username: $form.admin_username_field,
                    ipaddress: form_field_ipaddress,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                loader_msg = json.message
                EditData($form.admin_username_field)
                form_field_ipaddress = "";
            } else if (json.status == 403) {
                loader_msg = json.message
            } else {
                loader_msg = json.message;
            }
            buttonLoading_class = "btn btn-primary"
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
        } else {
            alert(msg_error);
        }
    }
    async function deleteIpList(e) {
        const res = await fetch(path_api+"api/deleteadminiplist", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idcompiplist: parseInt(e),
                username: $form.admin_username_field,
                page:"ADMIN-SAVE",
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            admin_listip = [];
            EditData($form.admin_username_field)
        }else if(json.status == 403){
            alert(json.message)
        }
    }
    async function statusCRM(e) {
        status_crm_satu = e
        if(e == "VALID"){
            isModalLoading = false;
            isModal_Form_New = true;
        }
    }
    async function statusCRMVALID(e) {
        switch(e){
            case "DEPOSIT":
                status_crm_dua = "DEPOSIT"
                img_deposit = img_check
                img_reject = "reject.svg"
                img_noanswer = "no-phone.svg"
                panel_deposit = true
                panel_reject = false
                panel_noanswer = false
                break;
            case "REJECT":
                status_crm_dua = "REJECT"
                img_deposit = "deposit.png"
                img_reject = img_check
                img_noanswer = "no-phone.svg"
                panel_deposit = false
                panel_reject = true
                panel_noanswer = false
                break;
            case "NOANSWER":
                status_crm_dua = "NOANSWER"
                img_deposit = "deposit.png"
                img_reject = "reject.svg"
                img_noanswer = img_check
                panel_deposit = false
                panel_reject = false
                panel_noanswer = true
                break;
        }
        
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const NewData = () => {
        sData = "New";
        isInput_username_enabled = true;
        clearField()
        isModal_Form_New = true;
    };
    const handleNewListIp = () => {
        isModal_Form_Listipaddress = true;
    }
   
    function clearField(){
        if(sData == "Edit"){
            admin_listrule = []
            admin_listip = []
        } 
        $form.admin_username_field = "";
        $form.admin_password_field = "";
        $form.admin_name_field = "";
        $form.admin_idrule_field = "0";
        $form.admin_status_field = "";
        form_field_ipaddress = "";
    }
    
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.admin_username
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) ||
                    item.admin_nama
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.admin_rule
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }

    
</script>
<Loader loader_class="{loader_class}" loader_msg="{loader_msg}" />

<Panel
    on:handleNewForm={NewData}
    on:handleRefresh={RefreshHalaman}
    panel_button_new={false}
    panel_button_refresh={true}
    panel_page="{page}"
    panel_total="{totalrecord}">
    <slot:template slot="panel_body">
        <div class="tabs tabs-boxed">
            <a class="tab tab-active">PROCESS</a> 
            <a class="tab">VALID</a> 
            <a class="tab">INVALID</a> 
        </div>
        <table class="table table-compact w-full mt-4">
            <thead class="sticky top-0">
                <tr>
                    <th width="7%" class="bg-[#475289] {font_size} text-white text-left">PHONE</th>
                    <th width="7%" class="bg-[#475289] {font_size} text-white text-left">WHATSAPP</th>
                    <th width="*" class="bg-[#475289] {font_size} text-white text-left">NAMA</th>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center"></th>
                </tr>
            </thead>
            {#if filterHome != ""}
                <tbody class="select-none">
                    {#each filterHome as rec}
                    <tr>
                        <td class="{font_size} align-top text-center">
                            <a class="text-blue-600 underline" href="tel:{rec.home_phone}">
                                {rec.home_phonealias}
                            </a>
                        </td>
                        <td class="{font_size} align-top text-left">
                            <a class="text-blue-600 underline" href="https://wa.me/{rec.home_phone}" target="_blank">
                                WHATSAPP
                            </a>
                        </td>
                        <td class="{font_size} align-top text-left">{rec.home_name}</td>
                        <td class="{font_size} align-top text-left">
                            <span
                                on:click={() => {
                                    statusCRM("VALID");
                                }} class="bg-[#ebfbee] text-[#6ec07b] text-center rounded-md p-1 px-2 cursor-pointer">VALID</span>
                            <span on:click={() => {
                                statusCRM("INVALID");
                            }} class="bg-[#fde3e3] text-[#ea7779] text-center rounded-md p-1 px-2 cursor-pointer">INVALID</span>
                        </td>
                    </tr>
                    {/each}
                </tbody>
            {:else}
                <tbody>
                    <tr>
                        <td colspan="10" class="text-center">
                            <progress class="self-start progress progress-primary w-56"></progress>
                        </td>
                    </tr>
                </tbody>
            {/if}
        </table>
    </slot:template>
</Panel>


<input type="checkbox" id="my-modal-formnew" class="modal-toggle" bind:checked={isModal_Form_New}>
<Modal_popup
    modal_popup_id="my-modal-formnew"
    modal_popup_title=""
    modal_popup_class="select-none max-w-full lg:max-w-xl overflow-hidden">
    <slot:template slot="modalpopup_body">
        <div class="flex flex-col w-full gap-2">
            <div class="flex justify-center gap-10 w-full">
                <div
                    on:click={() => {
                        statusCRMVALID("DEPOSIT");
                    }} 
                    class="flex flex-col justify-items-center text-center">
                    <div class="w-24 p-5 border-4 border-[#ffffff] bg-[#f1f1f1] rounded-full shadow-lg cursor-pointer">
                        <img width="60" class="" src="{img_deposit}" alt="">
                    </div>  
                    <span class="font-bold text-sm">DEPOSIT</span>
                </div>
                <div
                    on:click={() => {
                        statusCRMVALID("REJECT");
                    }} 
                    class="flex flex-col justify-items-center text-center">
                    <div class="w-24 p-5 border-4 border-[#ffffff] bg-[#f1f1f1] rounded-full shadow-lg cursor-pointer">
                        <img width="60" class="" src="{img_reject}" alt="">
                    </div>  
                    <span class="font-bold text-sm">REJECT</span>
                </div>
                <div
                    on:click={() => {
                        statusCRMVALID("NOANSWER");
                    }} 
                    class="flex flex-col justify-items-center text-center">
                    <div class="w-24 p-5 border-4 border-[#ffffff] bg-[#f1f1f1] rounded-full shadow-lg cursor-pointer">
                        <img width="60" class="" src="{img_noanswer}" alt="">
                    </div>  
                    <span class="font-bold text-sm">NO ANSWER</span>
                </div>
            </div>
            {#if panel_deposit}
                <div class="relative form-control mt-2">
                    <select
                        bind:value={website_field}
                        class="select select-bordered select-sm rounded-sm w-full focus:ring-0 focus:outline-none mb-2">
                        <option value="">--Pilih Website Agen--</option>
                        <option value="">ISB388</option>
                        <option value="">INDOSUPERBET</option>
                    </select>
                </div>
                <div class="relative form-control mt-2">
                    <Input_custom
                        input_enabled={true}
                        input_tipe="number"
                        bind:value={deposit_field}
                        input_id="deposit_field"
                        input_placeholder="Deposit"/>
                </div>
                <button class="btn btn-primary">Update</button>
            {/if}
            {#if panel_reject}
                <textarea class="textarea textarea-bordered" placeholder="Note"></textarea>
                <button class="btn btn-primary">Update</button>
            {/if}
            {#if panel_noanswer}
                <textarea class="textarea textarea-bordered" placeholder="Note"></textarea>
                <button class="btn btn-primary">Update</button>
            {/if}
        </div>
        
        
    </slot:template>
</Modal_popup>



<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />



