const utils = {

    getResponseObject: function () {
        let o = {
            success: false,
            msg: "",
            data: null
        }
        return Object.assign({}, o)
    },


}


module.exports = utils;