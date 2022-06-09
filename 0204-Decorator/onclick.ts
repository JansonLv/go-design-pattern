class Component1 {
    // 具有点击事件
    public onClick = () => {
        console.log('component1 click');
    };
}

class Decorator1 extends Component1 {
    private _component: Component1;

    constructor(component: Component1) {
        super();
        this._component = component;
    }

    public onClick = () => {
        console.log('打点');
        this._component.onClick();
    };
}

class Decorator2 extends Component1 {
    private _component: Component1;

    constructor(component: Component1) {
        super();
        this._component = component;
    }

    public onClick = () => {
        console.log('上报');
        this._component.onClick();
    };
}

const component1 = new Component1();
// 一个普通的点击
component1.onClick();
console.log();
console.log('wapper');
const wrapperComponent = new Decorator1(component1);
const reportComponent = new Decorator2(wrapperComponent);
// 一个具有打点功能的点击
reportComponent.onClick();
