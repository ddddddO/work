interface AddressFormat {
    zip: string;
    prefecture: string;
    city: string;

    generate: (value: string) => boolean;
 }

class Address implements AddressFormat {
    private addresses: any;
    public prefecture: string;
    public city: string;

    public constructor(private _zip: string) {
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
    };

    public getAddress(): string {
        let here = this.addresses[this._zip];
        return `${here.prefecture} ${here.city}`;
    };

    get zip(): string {
        return this._zip;
    };
    set zip(value: string) {
        this._zip = value;
    };

    public generate(value: string): boolean {
        console.log(value)
        return true
    }
}

let myaddress = new Address('038-0000');
console.log(myaddress.getAddress());
console.log(myaddress.zip)
myaddress.zip = 'xxxx'
console.log(myaddress.zip)

myaddress.generate('aaa')