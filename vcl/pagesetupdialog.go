
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

type TPageSetupDialog struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPageSetupDialog(owner IComponent) *TPageSetupDialog {
    p := new(TPageSetupDialog)
    p.instance = PageSetupDialog_Create(CheckPtr(owner))
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsPageSetupDialog(obj interface{}) *TPageSetupDialog {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TPageSetupDialog{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsPageSetupDialog.
func PageSetupDialogFromInst(inst uintptr) *TPageSetupDialog {
    return AsPageSetupDialog(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsPageSetupDialog.
func PageSetupDialogFromObj(obj IObject) *TPageSetupDialog {
    return AsPageSetupDialog(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsPageSetupDialog.
func PageSetupDialogFromUnsafePointer(ptr unsafe.Pointer) *TPageSetupDialog {
    return AsPageSetupDialog(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (p *TPageSetupDialog) Free() {
    if p.instance != 0 {
        PageSetupDialog_Free(p.instance)
        p.instance, p.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPageSetupDialog) Instance() uintptr {
    return p.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPageSetupDialog) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPageSetupDialog) IsValid() bool {
    return p.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (p *TPageSetupDialog) Is() TIs {
    return TIs(p.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (p *TPageSetupDialog) As() TAs {
//    return TAs(p.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPageSetupDialogClass() TClass {
    return PageSetupDialog_StaticClassType()
}

// CN: 执行。
// EN: .
func (p *TPageSetupDialog) Execute() bool {
    return PageSetupDialog_Execute(p.instance)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (p *TPageSetupDialog) FindComponent(AName string) *TComponent {
    return AsComponent(PageSetupDialog_FindComponent(p.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TPageSetupDialog) GetNamePath() string {
    return PageSetupDialog_GetNamePath(p.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (p *TPageSetupDialog) HasParent() bool {
    return PageSetupDialog_HasParent(p.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TPageSetupDialog) Assign(Source IObject) {
    PageSetupDialog_Assign(p.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPageSetupDialog) ClassType() TClass {
    return PageSetupDialog_ClassType(p.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPageSetupDialog) ClassName() string {
    return PageSetupDialog_ClassName(p.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPageSetupDialog) InstanceSize() int32 {
    return PageSetupDialog_InstanceSize(p.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPageSetupDialog) InheritsFrom(AClass TClass) bool {
    return PageSetupDialog_InheritsFrom(p.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPageSetupDialog) Equals(Obj IObject) bool {
    return PageSetupDialog_Equals(p.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPageSetupDialog) GetHashCode() int32 {
    return PageSetupDialog_GetHashCode(p.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (p *TPageSetupDialog) ToString() string {
    return PageSetupDialog_ToString(p.instance)
}

func (p *TPageSetupDialog) MarginLeft() int32 {
    return PageSetupDialog_GetMarginLeft(p.instance)
}

func (p *TPageSetupDialog) SetMarginLeft(value int32) {
    PageSetupDialog_SetMarginLeft(p.instance, value)
}

func (p *TPageSetupDialog) MarginTop() int32 {
    return PageSetupDialog_GetMarginTop(p.instance)
}

func (p *TPageSetupDialog) SetMarginTop(value int32) {
    PageSetupDialog_SetMarginTop(p.instance, value)
}

func (p *TPageSetupDialog) MarginRight() int32 {
    return PageSetupDialog_GetMarginRight(p.instance)
}

func (p *TPageSetupDialog) SetMarginRight(value int32) {
    PageSetupDialog_SetMarginRight(p.instance, value)
}

func (p *TPageSetupDialog) MarginBottom() int32 {
    return PageSetupDialog_GetMarginBottom(p.instance)
}

func (p *TPageSetupDialog) SetMarginBottom(value int32) {
    PageSetupDialog_SetMarginBottom(p.instance, value)
}

func (p *TPageSetupDialog) Options() TPageSetupDialogOptions {
    return PageSetupDialog_GetOptions(p.instance)
}

func (p *TPageSetupDialog) SetOptions(value TPageSetupDialogOptions) {
    PageSetupDialog_SetOptions(p.instance, value)
}

func (p *TPageSetupDialog) PageWidth() int32 {
    return PageSetupDialog_GetPageWidth(p.instance)
}

func (p *TPageSetupDialog) SetPageWidth(value int32) {
    PageSetupDialog_SetPageWidth(p.instance, value)
}

func (p *TPageSetupDialog) PageHeight() int32 {
    return PageSetupDialog_GetPageHeight(p.instance)
}

func (p *TPageSetupDialog) SetPageHeight(value int32) {
    PageSetupDialog_SetPageHeight(p.instance, value)
}

func (p *TPageSetupDialog) Units() TPageMeasureUnits {
    return PageSetupDialog_GetUnits(p.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (p *TPageSetupDialog) Handle() HWND {
    return PageSetupDialog_GetHandle(p.instance)
}

func (p *TPageSetupDialog) SetOnClose(fn TNotifyEvent) {
    PageSetupDialog_SetOnClose(p.instance, fn)
}

// CN: 设置显示事件。
// EN: .
func (p *TPageSetupDialog) SetOnShow(fn TNotifyEvent) {
    PageSetupDialog_SetOnShow(p.instance, fn)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (p *TPageSetupDialog) ComponentCount() int32 {
    return PageSetupDialog_GetComponentCount(p.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (p *TPageSetupDialog) ComponentIndex() int32 {
    return PageSetupDialog_GetComponentIndex(p.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (p *TPageSetupDialog) SetComponentIndex(value int32) {
    PageSetupDialog_SetComponentIndex(p.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (p *TPageSetupDialog) Owner() *TComponent {
    return AsComponent(PageSetupDialog_GetOwner(p.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (p *TPageSetupDialog) Name() string {
    return PageSetupDialog_GetName(p.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (p *TPageSetupDialog) SetName(value string) {
    PageSetupDialog_SetName(p.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (p *TPageSetupDialog) Tag() int {
    return PageSetupDialog_GetTag(p.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (p *TPageSetupDialog) SetTag(value int) {
    PageSetupDialog_SetTag(p.instance, value)
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (p *TPageSetupDialog) Components(AIndex int32) *TComponent {
    return AsComponent(PageSetupDialog_GetComponents(p.instance, AIndex))
}

