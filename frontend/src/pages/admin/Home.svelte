<script>
    import { createEventDispatcher } from "svelte";
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
    export let listwebsiteagen = [];
    export let totalrecord = 0;

    let page = "CRM PROCESS";
    let sData = "New";
    let isModal_Form_New = false
    let isModalLoading = false
    let isModalNotif = false
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_class = "btn btn-primary"
    let msg_error = "";

    let status_crm_satu = "";
    let status_crm_dua = "";
    let phone_field = "";
    let note_field = "";
    let idusersales_field = 0
    let idcrmsales_field = 0
    let website_field = "0"
    let iduseragen_field = ""
    let deposit_field = 0
    let img_deposit = "deposit.png"
    let img_reject = "reject.svg"
    let img_noanswer = "no-phone.svg"
    let img_check = "check.svg"
    let panel_deposit = false
    let panel_reject = false
    let panel_noanswer = false
    

  
    let searchHome = "";
    let filterHome = [];
    let dispatch = createEventDispatcher();
   
    async function SaveTransaksi() {
        let flag = true;
        msg_error = "";
        if(parseInt(idusersales_field) < 1){
            flag = true;
            msg_error += "The User Sales is required<br>";
        }
        if(parseInt(idcrmsales_field) < 1){
            flag = true;
            msg_error += "The CRM Sales is required<br>";
        }
        if(status_crm_satu == ""){
            flag = true;
            msg_error += "The Status is required<br>";
        }
        if (flag) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/salessave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page: "ADMIN-SAVE",
                    crm_idusersales: parseInt(idusersales_field),
                    crm_idcrmsales: parseInt(idcrmsales_field),
                    crm_idcwebagen: parseInt(website_field),
                    crm_status: status_crm_satu,
                    crm_status_dua: status_crm_dua,
                    crm_note: note_field,
                    crm_iduseragen: iduseragen_field,
                    crm_phone: phone_field,
                    crm_deposit: parseInt(deposit_field),
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
                    clearField()
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
            isModalLoading = false;
            isModal_Form_New = false;
        } else {
            alert(msg_error);
        }
    }
    
    async function statusCRM(e,idcrmsales,idusersales,phone) {
        resetstatus()
        status_crm_satu = e
        phone_field = phone;
        idusersales_field = parseInt(idusersales)
        idcrmsales_field = parseInt(idcrmsales)

        
        if(e == "VALID"){
            isModalLoading = false;
            isModal_Form_New = true;
        }else{
            SaveTransaksi()
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
        
    };
   
    function resetstatus(){
        status_crm_satu = ""
        phone_field = "";
        idusersales_field = 0
        idcrmsales_field = 0
        img_deposit = "deposit.png"
        img_reject = "reject.svg"
        img_noanswer = "no-phone.svg"
        panel_deposit = false
        panel_reject = false
        panel_noanswer = false
    }
    function clearField(){
        status_crm_satu = "";
        status_crm_dua = "";
        phone_field = "";
        note_field = "";
        idusersales_field = 0
        idcrmsales_field = 0
        website_field = "0"
        iduseragen_field = ""
        deposit_field = 0
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
                                    statusCRM("VALID",rec.home_idcrmsales,rec.home_idusersales,rec.home_phone);
                                }} class="bg-[#ebfbee] text-[#6ec07b] text-center rounded-md p-1 px-2 cursor-pointer">VALID</span>
                            <span on:click={() => {
                                statusCRM("INVALID",rec.home_idcrmsales,rec.home_idusersales,rec.home_phone);
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
                        <option value="0">--Pilih Website Agen--</option>
                        {#each listwebsiteagen as rec}
                            <option value="{rec.websiteagen_id}">{rec.websiteagen_name}</option>
                        {/each}
                    </select>
                </div>
                <div class="relative form-control mt-2">
                    <Input_custom
                        input_autofocus={false}
                        input_required={true}
                        input_tipe="text"
                        bind:value={iduseragen_field}
                        input_id="iduseragen_field"
                        input_placeholder="ID User Agen"/>
                </div>
                <div class="relative form-control mt-2">
                    <Input_custom
                        input_enabled={true}
                        input_tipe="number"
                        bind:value={deposit_field}
                        input_id="deposit_field"
                        input_placeholder="Deposit"/>
                </div>
                <button
                    on:click={() => {
                        SaveTransaksi();
                    }} 
                    class="btn btn-primary">Update</button>
            {/if}
            {#if panel_reject}
                <textarea 
                    bind:value={note_field}
                    class="textarea textarea-bordered" placeholder="Note"></textarea>
                <button on:click={() => {
                    SaveTransaksi();
                }} class="btn btn-primary">Update</button>
            {/if}
            {#if panel_noanswer}
                <textarea
                    bind:value={note_field} 
                    class="textarea textarea-bordered" placeholder="Note"></textarea>
                <button on:click={() => {
                    SaveTransaksi();
                }} class="btn btn-primary">Update</button>
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



