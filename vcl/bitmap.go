
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

type TBitmap struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewBitmap() *TBitmap {
    b := new(TBitmap)
    b.instance = Bitmap_Create()
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsBitmap(obj interface{}) *TBitmap {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TBitmap{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsBitmap.
func BitmapFromInst(inst uintptr) *TBitmap {
    return AsBitmap(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsBitmap.
func BitmapFromObj(obj IObject) *TBitmap {
    return AsBitmap(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsBitmap.
func BitmapFromUnsafePointer(ptr unsafe.Pointer) *TBitmap {
    return AsBitmap(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (b *TBitmap) Free() {
    if b.instance != 0 {
        Bitmap_Free(b.instance)
        b.instance, b.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TBitmap) Instance() uintptr {
    return b.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TBitmap) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TBitmap) IsValid() bool {
    return b.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (b *TBitmap) Is() TIs {
    return TIs(b.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (b *TBitmap) As() TAs {
//    return TAs(b.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TBitmapClass() TClass {
    return Bitmap_StaticClassType()
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (b *TBitmap) Assign(Source IObject) {
    Bitmap_Assign(b.instance, CheckPtr(Source))
}

func (b *TBitmap) FreeImage() {
    Bitmap_FreeImage(b.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (b *TBitmap) HandleAllocated() bool {
    return Bitmap_HandleAllocated(b.instance)
}

// CN: 文件流加载。
// EN: .
func (b *TBitmap) LoadFromStream(Stream IObject) {
    Bitmap_LoadFromStream(b.instance, CheckPtr(Stream))
}

// CN: 保存至流。
// EN: .
func (b *TBitmap) SaveToStream(Stream IObject) {
    Bitmap_SaveToStream(b.instance, CheckPtr(Stream))
}

func (b *TBitmap) SetSize(AWidth int32, AHeight int32) {
    Bitmap_SetSize(b.instance, AWidth , AHeight)
}

func (b *TBitmap) LoadFromResourceName(Instance uintptr, ResName string) {
    Bitmap_LoadFromResourceName(b.instance, Instance , ResName)
}

func (b *TBitmap) LoadFromResourceID(Instance uintptr, ResID int32) {
    Bitmap_LoadFromResourceID(b.instance, Instance , ResID)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TBitmap) Equals(Obj IObject) bool {
    return Bitmap_Equals(b.instance, CheckPtr(Obj))
}

// CN: 从文件加载。
// EN: .
func (b *TBitmap) LoadFromFile(Filename string) {
    Bitmap_LoadFromFile(b.instance, Filename)
}

// CN: 保存至文件。
// EN: .
func (b *TBitmap) SaveToFile(Filename string) {
    Bitmap_SaveToFile(b.instance, Filename)
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (b *TBitmap) GetNamePath() string {
    return Bitmap_GetNamePath(b.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TBitmap) ClassType() TClass {
    return Bitmap_ClassType(b.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TBitmap) ClassName() string {
    return Bitmap_ClassName(b.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TBitmap) InstanceSize() int32 {
    return Bitmap_InstanceSize(b.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TBitmap) InheritsFrom(AClass TClass) bool {
    return Bitmap_InheritsFrom(b.instance, AClass)
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TBitmap) GetHashCode() int32 {
    return Bitmap_GetHashCode(b.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (b *TBitmap) ToString() string {
    return Bitmap_ToString(b.instance)
}

// CN: 获取画布。
// EN: .
func (b *TBitmap) Canvas() *TCanvas {
    return AsCanvas(Bitmap_GetCanvas(b.instance))
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (b *TBitmap) Handle() HBITMAP {
    return Bitmap_GetHandle(b.instance)
}

// CN: 设置控件句柄。
// EN: Set Control handle.
func (b *TBitmap) SetHandle(value HBITMAP) {
    Bitmap_SetHandle(b.instance, value)
}

func (b *TBitmap) HandleType() TBitmapHandleType {
    return Bitmap_GetHandleType(b.instance)
}

func (b *TBitmap) SetHandleType(value TBitmapHandleType) {
    Bitmap_SetHandleType(b.instance, value)
}

func (b *TBitmap) MaskHandle() HBITMAP {
    return Bitmap_GetMaskHandle(b.instance)
}

func (b *TBitmap) SetMaskHandle(value HBITMAP) {
    Bitmap_SetMaskHandle(b.instance, value)
}

func (b *TBitmap) PixelFormat() TPixelFormat {
    return Bitmap_GetPixelFormat(b.instance)
}

func (b *TBitmap) SetPixelFormat(value TPixelFormat) {
    Bitmap_SetPixelFormat(b.instance, value)
}

func (b *TBitmap) TransparentMode() TTransparentMode {
    return Bitmap_GetTransparentMode(b.instance)
}

func (b *TBitmap) SetTransparentMode(value TTransparentMode) {
    Bitmap_SetTransparentMode(b.instance, value)
}

func (b *TBitmap) Empty() bool {
    return Bitmap_GetEmpty(b.instance)
}

// CN: 获取高度。
// EN: Get height.
func (b *TBitmap) Height() int32 {
    return Bitmap_GetHeight(b.instance)
}

// CN: 设置高度。
// EN: Set height.
func (b *TBitmap) SetHeight(value int32) {
    Bitmap_SetHeight(b.instance, value)
}

// CN: 获取修改。
// EN: Get modified.
func (b *TBitmap) Modified() bool {
    return Bitmap_GetModified(b.instance)
}

// CN: 设置修改。
// EN: Set modified.
func (b *TBitmap) SetModified(value bool) {
    Bitmap_SetModified(b.instance, value)
}

func (b *TBitmap) Palette() HPALETTE {
    return Bitmap_GetPalette(b.instance)
}

func (b *TBitmap) SetPalette(value HPALETTE) {
    Bitmap_SetPalette(b.instance, value)
}

func (b *TBitmap) PaletteModified() bool {
    return Bitmap_GetPaletteModified(b.instance)
}

func (b *TBitmap) SetPaletteModified(value bool) {
    Bitmap_SetPaletteModified(b.instance, value)
}

// CN: 获取透明。
// EN: Get transparent.
func (b *TBitmap) Transparent() bool {
    return Bitmap_GetTransparent(b.instance)
}

// CN: 设置透明。
// EN: Set transparent.
func (b *TBitmap) SetTransparent(value bool) {
    Bitmap_SetTransparent(b.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (b *TBitmap) Width() int32 {
    return Bitmap_GetWidth(b.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (b *TBitmap) SetWidth(value int32) {
    Bitmap_SetWidth(b.instance, value)
}

// CN: 设置改变事件。
// EN: Set changed event.
func (b *TBitmap) SetOnChange(fn TNotifyEvent) {
    Bitmap_SetOnChange(b.instance, fn)
}

func (b *TBitmap) ScanLine(Row int32) uintptr {
    return Bitmap_GetScanLine(b.instance, Row)
}

