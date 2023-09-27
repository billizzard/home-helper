class SkinnedMesh extends THREE.Mesh {
    constructor(geometry, materials) {
        super(geometry, materials);

        this.idMatrix = SkinnedMesh.defaultMatrix();
        this.bones = [];
        this.boneMatrices = [];
        //...
    }
    update(camera) {
        //...
        super.update();
    }
    static defaultMatrix() {
        return new THREE.Matrix4();
    }
}

const obj = {a:1, b:2, c:3, d:4, e:5 };
const result = (({a,b,c}) => ({a, b, c}))(obj);
console.log(result);