export default {
    state: {
        userInfo: {},
        appSheetList: [],
        planeList: [],
        trainList:[],
        planeInfo: [],
        planeSearch: {},

        //酒店相关全局对象
        hotelList:[],
        hotelSearch: {},
        trainOrder:{},

    },
    getters: {
        getAppSheetListByAppId: (state) => (appId) => {
            return state.appSheetList.find((val) => {
                return val.appId === appId
            })
        },
    },
    mutations: {
        login(state, userInfo) {
            //TODO 因为会被存在sessionStorage中，需要加密
            state.userInfo = userInfo;
        },
        recordAppSheetList(state, appSheet) {
            //申请单 例子
            // let appSheet = {
            //     appId: "20191203121534123",
            //     appReason: "差旅事由",
            //     appProject: "(286)会计核算中心--(110000000)系公用经费 ",
            //     travelType: "科研", //学术，公务
            //     travelUsers: [
            //         {
            //             username: "renhaohotelResultList",
            //             usercode: "20080001",
            //             operator: true, //经办人
            //             travelLevel: "A3", //差旅标准
            //         },
            //         {
            //             username: "renhao",
            //             usercode: "20080001",
            //             travelLevel: "A3"
            //         },
            //         {
            //             username: "renhao",
            //             usercode: "20080001",
            //             travelLevel: "A3"
            //         },
            //     ],
            // };
            state.appSheetList.push(appSheet);
        },

        //预下单存储
        order(state, resultList){
            state.order = resultList;
        },
        //订单中心选择的订单类型
        orderType(state, type) {
            state.orderType = type;
        },

        //机票信息存储
        planeResultList(state, resultList) {
            state.planeList.push(...resultList);
        },
        //机票信息置空
        resetPlaneResultList(state) {
            state.planeList=[];
        },
        //机票预订信息
        bookPlaneInfo(state, pInfo) {
            state.planeInfo = pInfo;
        },
        //机票检索条件
        planeSearch(state, pConditon) {
            state.planeSearch = pConditon;
        },
        //机票预下单信息存储
        flightPreOrder(state, result) {
            state.flightPreOrder = result;
        },
        //原机票信息存储
        originalFlightOrder(state, result) {
            state.originalFlightOrder = result;
        },
        flightRefundInfo(state, flightRefundInfo) {
            state.flightRefundInfo = flightRefundInfo;
        },

        //火车票信息存储
        trainResultList(state, resultList) {
            state.trainList.push(...resultList);
        },
        //火车票列表置空
        resetTrainResultList(state) {
            state.trainList=[];
        },

        //选中的火车车次信息存储
        bookTrainInfo(state, resultList) {
            state.bookTrainInfo = resultList;
        },
        //选中的火车坐席信息存储
        bookTrainTicketInfo(state, resultList) {
            state.bookTrainTicketInfo = resultList;
        },
        //火车票原订单
        oldTrainOrder(state, oldTrainOrder){
            state.oldTrainOrder = oldTrainOrder;
        },
        //火车预下单存储
        trainOrder(state, resultList){
            state.trainOrder = resultList;
        },
        //火车改签预下单存储
        preChangeTrainOrder(state, result){
            state.preChangeTrainOrder = result;
        },

        //酒店信息存储
        hotelResultList(state, resultList) {
            state.hotelList.push(...resultList);
        },
        //酒店列表置空
        resetHotelResultList(state) {
            state.hotelList=[];
        },
        hotelSearch(state, pConditon) {
            state.hotelSearch = pConditon;
        },
        //酒店Item
        hotelItem(state, result) {
            state.hotelItem = result;
        },
        //酒店详情
        hotelDetail(state, result) {
            state.hotelDetail = result;
        },
        //酒店房间详情
        hotelRoomDetail(state, result) {
            state.hotelRoomDetail = result;
        },
        //酒店入住人
        customers(state, result) {
            state.customers = result;
        },
        //机票保险信息
        insuranceProduct(state, result) {
            state.insuranceProduct = result;
        },
    },
    actions: {}
}
