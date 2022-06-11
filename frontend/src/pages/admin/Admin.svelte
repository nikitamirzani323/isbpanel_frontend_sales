<script>
    import Home from "../admin/Home.svelte";
    export let path_api = ""
    export let font_size = "";
    let listHome = [];
    let record = "";
    let totalrecord = 0;
    let token = localStorage.getItem("token");
    let akses_page = true;
    let admin_listrule = [];
  
    async function initHome() {
        const res = await fetch(path_api+"api/sales", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            totalrecord = record.length;
            let no = 0
            let recordlistrule = json.listruleadmin;
            let status_class = "";
            let temp_phone = "";
            let temp_xxx = 0;
            let temp_alias = ""
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    if (record[i]["admin_status"] == "ACTIVE") {
                        status_class = "bg-[#ebfbee] text-[#6ec07b]"
                    } else {
                        status_class = "bg-[#fde3e3] text-[#ea7779]"
                    }
                    temp_phone = record[i]["crm_phone"];
                    temp_xxx = temp_phone.length - 5;
                    temp_alias = ""
                    for(let z=0;z<temp_phone.length;z++){
                        if(temp_xxx <= z){
                            temp_alias += "x"
                        }else{
                            temp_alias += temp_phone[z]
                        }
                        
                    }
                    
                    no = no + 1
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_idcrmsales: record[i]["crm_idcrmsales"],
                            home_idusersales: record[i]["crm_idusersales"],
                            home_name: record[i]["crm_name"],
                            home_phone: record[i]["crm_phone"],
                            home_phonealias: temp_alias,
                        },
                    ];
                }
            }
            if (recordlistrule != null) {
                for (let i = 0; i < recordlistrule.length; i++) {
                    admin_listrule = [
                        ...admin_listrule,
                        {
                            adminrule_idruleadmin:
                                recordlistrule[i]["adminrule_idruleadmin"],
                            adminrule_name: recordlistrule[i]["adminrule_name"],
                        },
                    ];
                }
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        admin_listrule = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 1000);
    };
    const handleLogout = (e) => {
        logout()
    }
    initHome();
</script>
{#if akses_page == true}
    <Home
        on:handleRefreshData={handleRefreshData}
        on:handleLogout={handleLogout}
        {path_api}
        {font_size}
        {token}
        {admin_listrule}
        {listHome}
        {totalrecord}/>
{/if}