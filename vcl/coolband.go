
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TCoolBand struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCoolBand(AOwner *TCollection) *TCoolBand {
    c := new(TCoolBand)
    c.instance = CoolBand_Create(CheckPtr(AOwner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsCoolBand(obj interface{}) *TCoolBand {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TCoolBand{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsCoolBand.
func CoolBandFromInst(inst uintptr) *TCoolBand {
    return AsCoolBand(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsCoolBand.
func CoolBandFromObj(obj IObject) *TCoolBand {
    return AsCoolBand(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsCoolBand.
func CoolBandFromUnsafePointer(ptr unsafe.Pointer) *TCoolBand {
    return AsCoolBand(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (c *TCoolBand) Free() {
    if c.instance != 0 {
        CoolBand_Free(c.instance)
        c.instance, c.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCoolBand) Instance() uintptr {
    return c.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCoolBand) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCoolBand) IsValid() bool {
    return c.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (c *TCoolBand) Is() TIs {
    return TIs(c.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (c *TCoolBand) As() TAs {
//    return TAs(c.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCoolBandClass() TClass {
    return CoolBand_StaticClassType()
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TCoolBand) Assign(Source IObject) {
    CoolBand_Assign(c.instance, CheckPtr(Source))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TCoolBand) GetNamePath() string {
    return CoolBand_GetNamePath(c.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCoolBand) ClassType() TClass {
    return CoolBand_ClassType(c.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCoolBand) ClassName() string {
    return CoolBand_ClassName(c.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCoolBand) InstanceSize() int32 {
    return CoolBand_InstanceSize(c.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCoolBand) InheritsFrom(AClass TClass) bool {
    return CoolBand_InheritsFrom(c.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCoolBand) Equals(Obj IObject) bool {
    return CoolBand_Equals(c.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCoolBand) GetHashCode() int32 {
    return CoolBand_GetHashCode(c.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (c *TCoolBand) ToString() string {
    return CoolBand_ToString(c.instance)
}

// CN: 获取高度。
// EN: Get height.
func (c *TCoolBand) Height() int32 {
    return CoolBand_GetHeight(c.instance)
}

func (c *TCoolBand) Bitmap() *TBitmap {
    return AsBitmap(CoolBand_GetBitmap(c.instance))
}

func (c *TCoolBand) SetBitmap(value *TBitmap) {
    CoolBand_SetBitmap(c.instance, CheckPtr(value))
}

// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (c *TCoolBand) BorderStyle() TBorderStyle {
    return CoolBand_GetBorderStyle(c.instance)
}

// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (c *TCoolBand) SetBorderStyle(value TBorderStyle) {
    CoolBand_SetBorderStyle(c.instance, value)
}

func (c *TCoolBand) Break() bool {
    return CoolBand_GetBreak(c.instance)
}

func (c *TCoolBand) SetBreak(value bool) {
    CoolBand_SetBreak(c.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (c *TCoolBand) Color() TColor {
    return CoolBand_GetColor(c.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (c *TCoolBand) SetColor(value TColor) {
    CoolBand_SetColor(c.instance, value)
}

func (c *TCoolBand) Control() *TWinControl {
    return AsWinControl(CoolBand_GetControl(c.instance))
}

func (c *TCoolBand) SetControl(value IWinControl) {
    CoolBand_SetControl(c.instance, CheckPtr(value))
}

func (c *TCoolBand) FixedBackground() bool {
    return CoolBand_GetFixedBackground(c.instance)
}

func (c *TCoolBand) SetFixedBackground(value bool) {
    CoolBand_SetFixedBackground(c.instance, value)
}

func (c *TCoolBand) FixedSize() bool {
    return CoolBand_GetFixedSize(c.instance)
}

func (c *TCoolBand) SetFixedSize(value bool) {
    CoolBand_SetFixedSize(c.instance, value)
}

func (c *TCoolBand) HorizontalOnly() bool {
    return CoolBand_GetHorizontalOnly(c.instance)
}

func (c *TCoolBand) SetHorizontalOnly(value bool) {
    CoolBand_SetHorizontalOnly(c.instance, value)
}

// CN: 获取图像在images中的索引。
// EN: .
func (c *TCoolBand) ImageIndex() int32 {
    return CoolBand_GetImageIndex(c.instance)
}

// CN: 设置图像在images中的索引。
// EN: .
func (c *TCoolBand) SetImageIndex(value int32) {
    CoolBand_SetImageIndex(c.instance, value)
}

func (c *TCoolBand) MinHeight() int32 {
    return CoolBand_GetMinHeight(c.instance)
}

func (c *TCoolBand) SetMinHeight(value int32) {
    CoolBand_SetMinHeight(c.instance, value)
}

func (c *TCoolBand) MinWidth() int32 {
    return CoolBand_GetMinWidth(c.instance)
}

func (c *TCoolBand) SetMinWidth(value int32) {
    CoolBand_SetMinWidth(c.instance, value)
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (c *TCoolBand) ParentColor() bool {
    return CoolBand_GetParentColor(c.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (c *TCoolBand) SetParentColor(value bool) {
    CoolBand_SetParentColor(c.instance, value)
}

func (c *TCoolBand) ParentBitmap() bool {
    return CoolBand_GetParentBitmap(c.instance)
}

func (c *TCoolBand) SetParentBitmap(value bool) {
    CoolBand_SetParentBitmap(c.instance, value)
}

// CN: 获取文本。
// EN: .
func (c *TCoolBand) Text() string {
    return CoolBand_GetText(c.instance)
}

// CN: 设置文本。
// EN: .
func (c *TCoolBand) SetText(value string) {
    CoolBand_SetText(c.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TCoolBand) Visible() bool {
    return CoolBand_GetVisible(c.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TCoolBand) SetVisible(value bool) {
    CoolBand_SetVisible(c.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (c *TCoolBand) Width() int32 {
    return CoolBand_GetWidth(c.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (c *TCoolBand) SetWidth(value int32) {
    CoolBand_SetWidth(c.instance, value)
}

func (c *TCoolBand) Collection() *TCollection {
    return AsCollection(CoolBand_GetCollection(c.instance))
}

func (c *TCoolBand) SetCollection(value *TCollection) {
    CoolBand_SetCollection(c.instance, CheckPtr(value))
}

func (c *TCoolBand) Index() int32 {
    return CoolBand_GetIndex(c.instance)
}

func (c *TCoolBand) SetIndex(value int32) {
    CoolBand_SetIndex(c.instance, value)
}

func (c *TCoolBand) DisplayName() string {
    return CoolBand_GetDisplayName(c.instance)
}

func (c *TCoolBand) SetDisplayName(value string) {
    CoolBand_SetDisplayName(c.instance, value)
}

