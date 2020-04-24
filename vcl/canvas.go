
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

type TCanvas struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCanvas() *TCanvas {
    c := new(TCanvas)
    c.instance = Canvas_Create()
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsCanvas(obj interface{}) *TCanvas {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TCanvas{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsCanvas.
func CanvasFromInst(inst uintptr) *TCanvas {
    return AsCanvas(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsCanvas.
func CanvasFromObj(obj IObject) *TCanvas {
    return AsCanvas(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsCanvas.
func CanvasFromUnsafePointer(ptr unsafe.Pointer) *TCanvas {
    return AsCanvas(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (c *TCanvas) Free() {
    if c.instance != 0 {
        Canvas_Free(c.instance)
        c.instance, c.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCanvas) Instance() uintptr {
    return c.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCanvas) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCanvas) IsValid() bool {
    return c.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (c *TCanvas) Is() TIs {
    return TIs(c.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (c *TCanvas) As() TAs {
//    return TAs(c.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCanvasClass() TClass {
    return Canvas_StaticClassType()
}

func (c *TCanvas) Arc(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_Arc(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

func (c *TCanvas) ArcTo(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_ArcTo(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

func (c *TCanvas) AngleArc(X int32, Y int32, Radius uint32, StartAngle float32, SweepAngle float32) {
    Canvas_AngleArc(c.instance, X , Y , Radius , StartAngle , SweepAngle)
}

func (c *TCanvas) Chord(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_Chord(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

func (c *TCanvas) Ellipse(X1 int32, Y1 int32, X2 int32, Y2 int32) {
    Canvas_Ellipse(c.instance, X1 , Y1 , X2 , Y2)
}

func (c *TCanvas) FloodFill(X int32, Y int32, Color TColor, FillStyle TFillStyle) {
    Canvas_FloodFill(c.instance, X , Y , Color , FillStyle)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (c *TCanvas) HandleAllocated() bool {
    return Canvas_HandleAllocated(c.instance)
}

func (c *TCanvas) LineTo(X int32, Y int32) {
    Canvas_LineTo(c.instance, X , Y)
}

func (c *TCanvas) MoveTo(X int32, Y int32) {
    Canvas_MoveTo(c.instance, X , Y)
}

func (c *TCanvas) Pie(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_Pie(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

func (c *TCanvas) Rectangle(X1 int32, Y1 int32, X2 int32, Y2 int32) {
    Canvas_Rectangle(c.instance, X1 , Y1 , X2 , Y2)
}

// CN: 刷新控件。
// EN: Refresh control.
func (c *TCanvas) Refresh() {
    Canvas_Refresh(c.instance)
}

func (c *TCanvas) RoundRect(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32) {
    Canvas_RoundRect(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3)
}

func (c *TCanvas) TextExtent(Text string) TSize {
    return Canvas_TextExtent(c.instance, Text)
}

func (c *TCanvas) TextOut(X int32, Y int32, Text string) {
    Canvas_TextOut(c.instance, X , Y , Text)
}

func (c *TCanvas) Lock() {
    Canvas_Lock(c.instance)
}

func (c *TCanvas) TextHeight(Text string) int32 {
    return Canvas_TextHeight(c.instance, Text)
}

func (c *TCanvas) TextWidth(Text string) int32 {
    return Canvas_TextWidth(c.instance, Text)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TCanvas) Assign(Source IObject) {
    Canvas_Assign(c.instance, CheckPtr(Source))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TCanvas) GetNamePath() string {
    return Canvas_GetNamePath(c.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCanvas) ClassType() TClass {
    return Canvas_ClassType(c.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCanvas) ClassName() string {
    return Canvas_ClassName(c.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCanvas) InstanceSize() int32 {
    return Canvas_InstanceSize(c.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCanvas) InheritsFrom(AClass TClass) bool {
    return Canvas_InheritsFrom(c.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCanvas) Equals(Obj IObject) bool {
    return Canvas_Equals(c.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCanvas) GetHashCode() int32 {
    return Canvas_GetHashCode(c.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (c *TCanvas) ToString() string {
    return Canvas_ToString(c.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (c *TCanvas) Handle() HDC {
    return Canvas_GetHandle(c.instance)
}

// CN: 设置控件句柄。
// EN: Set Control handle.
func (c *TCanvas) SetHandle(value HDC) {
    Canvas_SetHandle(c.instance, value)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (c *TCanvas) Brush() *TBrush {
    return AsBrush(Canvas_GetBrush(c.instance))
}

// CN: 设置画刷对象。
// EN: Set Brush.
func (c *TCanvas) SetBrush(value *TBrush) {
    Canvas_SetBrush(c.instance, CheckPtr(value))
}

func (c *TCanvas) CopyMode() int32 {
    return Canvas_GetCopyMode(c.instance)
}

func (c *TCanvas) SetCopyMode(value int32) {
    Canvas_SetCopyMode(c.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (c *TCanvas) Font() *TFont {
    return AsFont(Canvas_GetFont(c.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (c *TCanvas) SetFont(value *TFont) {
    Canvas_SetFont(c.instance, CheckPtr(value))
}

func (c *TCanvas) Pen() *TPen {
    return AsPen(Canvas_GetPen(c.instance))
}

func (c *TCanvas) SetPen(value *TPen) {
    Canvas_SetPen(c.instance, CheckPtr(value))
}

// CN: 设置改变事件。
// EN: Set changed event.
func (c *TCanvas) SetOnChange(fn TNotifyEvent) {
    Canvas_SetOnChange(c.instance, fn)
}

func (c *TCanvas) SetOnChanging(fn TNotifyEvent) {
    Canvas_SetOnChanging(c.instance, fn)
}

