<script>
    import Home from "./Home.svelte";
    export let path_api = ""
    export let font_size = "";
    let listHome = [];
    let listwebsiteagen = [];
    let record = "";
    let totalrecord = 0;
    let token = localStorage.getItem("token");
    let akses_page = true;
  
    async function initHome(e) {
        const res = await fetch(path_api+"api/sales", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                crm_status:e
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            totalrecord = record.length;
            let no = 0
            let temp_phone = "";
            let temp_xxx = 0;
            let temp_alias = ""
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
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
        } else {
            logout();
        }
    }
    async function initWebsiteAgen() {
        listwebsiteagen = [];
        const res = await fetch(path_api+"api/websiteagen", {
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
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listwebsiteagen = [
                        ...listwebsiteagen,
                        {
                            websiteagen_id: record[i]["websiteagen_id"],
                            websiteagen_name: record[i]["websiteagen_name"],
                        },
                    ];
                }
            }
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        listwebsiteagen = [];
        totalrecord = 0;
        console.log(e.detail.tab_status)
        setTimeout(function () {
            initHome(e.detail.tab_status);
            initWebsiteAgen();
        }, 1000);
    };
    const handleLogout = (e) => {
        logout()
    }
    initHome("PROCESS");
    initWebsiteAgen();
</script>
{#if akses_page == true}
    <Home
        on:handleRefreshData={handleRefreshData}
        on:handleLogout={handleLogout}
        {path_api}
        {font_size}
        {token}
        {listHome}
        {listwebsiteagen}
        {totalrecord}/>
{/if}