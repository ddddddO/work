var Address = /** @class */ (function () {
    function Address(_zip) {
        this._zip = _zip;
        this.addresses =
            {
                "079-1100": {
                    "prefecture": "北海道",
                    "city": "赤平市"
                },
                "038-0000": {
                    "prefecture": "青森県",
                    "city": "青森市"
                }
            };
    }
    ;
    Address.prototype.getAddress = function () {
        var here = this.addresses[this._zip];
        return "".concat(here.prefecture, " ").concat(here.city);
    };
    ;
    Object.defineProperty(Address.prototype, "zip", {
        get: function () {
            return this._zip;
        },
        set: function (value) {
            this._zip = value;
        },
        enumerable: false,
        configurable: true
    });
    ;
    ;
    Address.prototype.generate = function (value) {
        console.log(value);
        return true;
    };
    return Address;
}());
var myaddress = new Address('038-0000');
console.log(myaddress.getAddress());
console.log(myaddress.zip);
myaddress.zip = 'xxxx';
console.log(myaddress.zip);
myaddress.generate('aaa');
