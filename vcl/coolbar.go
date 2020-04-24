
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

type TCoolBar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCoolBar(owner IComponent) *TCoolBar {
    c := new(TCoolBar)
    c.instance = CoolBar_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsCoolBar(obj interface{}) *TCoolBar {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TCoolBar{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsCoolBar.
func CoolBarFromInst(inst uintptr) *TCoolBar {
    return AsCoolBar(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsCoolBar.
func CoolBarFromObj(obj IObject) *TCoolBar {
    return AsCoolBar(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsCoolBar.
func CoolBarFromUnsafePointer(ptr unsafe.Pointer) *TCoolBar {
    return AsCoolBar(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (c *TCoolBar) Free() {
    if c.instance != 0 {
        CoolBar_Free(c.instance)
        c.instance, c.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCoolBar) Instance() uintptr {
    return c.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCoolBar) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCoolBar) IsValid() bool {
    return c.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (c *TCoolBar) Is() TIs {
    return TIs(c.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (c *TCoolBar) As() TAs {
//    return TAs(c.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCoolBarClass() TClass {
    return CoolBar_StaticClassType()
}

func (c *TCoolBar) FlipChildren(AllLevels bool) {
    CoolBar_FlipChildren(c.instance, AllLevels)
}

// CN: 是否可以获得焦点。
// EN: .
func (c *TCoolBar) CanFocus() bool {
    return CoolBar_CanFocus(c.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (c *TCoolBar) ContainsControl(Control IControl) bool {
    return CoolBar_ContainsControl(c.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (c *TCoolBar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(CoolBar_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (c *TCoolBar) DisableAlign() {
    CoolBar_DisableAlign(c.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (c *TCoolBar) EnableAlign() {
    CoolBar_EnableAlign(c.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (c *TCoolBar) FindChildControl(ControlName string) *TControl {
    return AsControl(CoolBar_FindChildControl(c.instance, ControlName))
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (c *TCoolBar) Focused() bool {
    return CoolBar_Focused(c.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (c *TCoolBar) HandleAllocated() bool {
    return CoolBar_HandleAllocated(c.instance)
}

// CN: 要求重绘。
// EN: Redraw.
func (c *TCoolBar) Invalidate() {
    CoolBar_Invalidate(c.instance)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (c *TCoolBar) RemoveControl(AControl IControl) {
    CoolBar_RemoveControl(c.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (c *TCoolBar) Realign() {
    CoolBar_Realign(c.instance)
}

// CN: 重绘。
// EN: Repaint.
func (c *TCoolBar) Repaint() {
    CoolBar_Repaint(c.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (c *TCoolBar) ScaleBy(M int32, D int32) {
    CoolBar_ScaleBy(c.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (c *TCoolBar) ScrollBy(DeltaX int32, DeltaY int32) {
    CoolBar_ScrollBy(c.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (c *TCoolBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CoolBar_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (c *TCoolBar) SetFocus() {
    CoolBar_SetFocus(c.instance)
}

// CN: 控件更新。
// EN: Update.
func (c *TCoolBar) Update() {
    CoolBar_Update(c.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (c *TCoolBar) BringToFront() {
    CoolBar_BringToFront(c.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (c *TCoolBar) ClientToScreen(Point TPoint) TPoint {
    return CoolBar_ClientToScreen(c.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (c *TCoolBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return CoolBar_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (c *TCoolBar) Dragging() bool {
    return CoolBar_Dragging(c.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (c *TCoolBar) HasParent() bool {
    return CoolBar_HasParent(c.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (c *TCoolBar) Hide() {
    CoolBar_Hide(c.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (c *TCoolBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CoolBar_Perform(c.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (c *TCoolBar) Refresh() {
    CoolBar_Refresh(c.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (c *TCoolBar) ScreenToClient(Point TPoint) TPoint {
    return CoolBar_ScreenToClient(c.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (c *TCoolBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return CoolBar_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (c *TCoolBar) SendToBack() {
    CoolBar_SendToBack(c.instance)
}

// CN: 显示控件。
// EN: Show control.
func (c *TCoolBar) Show() {
    CoolBar_Show(c.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (c *TCoolBar) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return CoolBar_GetTextBuf(c.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (c *TCoolBar) GetTextLen() int32 {
    return CoolBar_GetTextLen(c.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (c *TCoolBar) SetTextBuf(Buffer string) {
    CoolBar_SetTextBuf(c.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (c *TCoolBar) FindComponent(AName string) *TComponent {
    return AsComponent(CoolBar_FindComponent(c.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TCoolBar) GetNamePath() string {
    return CoolBar_GetNamePath(c.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TCoolBar) Assign(Source IObject) {
    CoolBar_Assign(c.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCoolBar) ClassType() TClass {
    return CoolBar_ClassType(c.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCoolBar) ClassName() string {
    return CoolBar_ClassName(c.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCoolBar) InstanceSize() int32 {
    return CoolBar_InstanceSize(c.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCoolBar) InheritsFrom(AClass TClass) bool {
    return CoolBar_InheritsFrom(c.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCoolBar) Equals(Obj IObject) bool {
    return CoolBar_Equals(c.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCoolBar) GetHashCode() int32 {
    return CoolBar_GetHashCode(c.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (c *TCoolBar) ToString() string {
    return CoolBar_ToString(c.instance)
}

func (c *TCoolBar) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    CoolBar_AnchorToNeighbour(c.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (c *TCoolBar) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    CoolBar_AnchorParallel(c.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (c *TCoolBar) AnchorHorizontalCenterTo(ASibling IControl) {
    CoolBar_AnchorHorizontalCenterTo(c.instance, CheckPtr(ASibling))
}

func (c *TCoolBar) AnchorVerticalCenterTo(ASibling IControl) {
    CoolBar_AnchorVerticalCenterTo(c.instance, CheckPtr(ASibling))
}

func (c *TCoolBar) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    CoolBar_AnchorAsAlign(c.instance, ATheAlign , ASpace)
}

func (c *TCoolBar) AnchorClient(ASpace int32) {
    CoolBar_AnchorClient(c.instance, ASpace)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (c *TCoolBar) Align() TAlign {
    return CoolBar_GetAlign(c.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (c *TCoolBar) SetAlign(value TAlign) {
    CoolBar_SetAlign(c.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (c *TCoolBar) Anchors() TAnchors {
    return CoolBar_GetAnchors(c.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (c *TCoolBar) SetAnchors(value TAnchors) {
    CoolBar_SetAnchors(c.instance, value)
}

// CN: 获取自动调整大小。
// EN: .
func (c *TCoolBar) AutoSize() bool {
    return CoolBar_GetAutoSize(c.instance)
}

// CN: 设置自动调整大小。
// EN: .
func (c *TCoolBar) SetAutoSize(value bool) {
    CoolBar_SetAutoSize(c.instance, value)
}

func (c *TCoolBar) BandBorderStyle() TBorderStyle {
    return CoolBar_GetBandBorderStyle(c.instance)
}

func (c *TCoolBar) SetBandBorderStyle(value TBorderStyle) {
    CoolBar_SetBandBorderStyle(c.instance, value)
}

func (c *TCoolBar) BandMaximize() TCoolBandMaximize {
    return CoolBar_GetBandMaximize(c.instance)
}

func (c *TCoolBar) SetBandMaximize(value TCoolBandMaximize) {
    CoolBar_SetBandMaximize(c.instance, value)
}

func (c *TCoolBar) Bands() *TCoolBands {
    return AsCoolBands(CoolBar_GetBands(c.instance))
}

func (c *TCoolBar) SetBands(value *TCoolBands) {
    CoolBar_SetBands(c.instance, CheckPtr(value))
}

// CN: 获取边框的宽度。
// EN: .
func (c *TCoolBar) BorderWidth() int32 {
    return CoolBar_GetBorderWidth(c.instance)
}

// CN: 设置边框的宽度。
// EN: .
func (c *TCoolBar) SetBorderWidth(value int32) {
    CoolBar_SetBorderWidth(c.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (c *TCoolBar) Color() TColor {
    return CoolBar_GetColor(c.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (c *TCoolBar) SetColor(value TColor) {
    CoolBar_SetColor(c.instance, value)
}

func (c *TCoolBar) Constraints() *TSizeConstraints {
    return AsSizeConstraints(CoolBar_GetConstraints(c.instance))
}

func (c *TCoolBar) SetConstraints(value *TSizeConstraints) {
    CoolBar_SetConstraints(c.instance, CheckPtr(value))
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (c *TCoolBar) DockSite() bool {
    return CoolBar_GetDockSite(c.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (c *TCoolBar) SetDockSite(value bool) {
    CoolBar_SetDockSite(c.instance, value)
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (c *TCoolBar) DoubleBuffered() bool {
    return CoolBar_GetDoubleBuffered(c.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (c *TCoolBar) SetDoubleBuffered(value bool) {
    CoolBar_SetDoubleBuffered(c.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (c *TCoolBar) DragCursor() TCursor {
    return CoolBar_GetDragCursor(c.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (c *TCoolBar) SetDragCursor(value TCursor) {
    CoolBar_SetDragCursor(c.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (c *TCoolBar) DragKind() TDragKind {
    return CoolBar_GetDragKind(c.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (c *TCoolBar) SetDragKind(value TDragKind) {
    CoolBar_SetDragKind(c.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (c *TCoolBar) DragMode() TDragMode {
    return CoolBar_GetDragMode(c.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (c *TCoolBar) SetDragMode(value TDragMode) {
    CoolBar_SetDragMode(c.instance, value)
}

func (c *TCoolBar) EdgeBorders() TEdgeBorders {
    return CoolBar_GetEdgeBorders(c.instance)
}

func (c *TCoolBar) SetEdgeBorders(value TEdgeBorders) {
    CoolBar_SetEdgeBorders(c.instance, value)
}

func (c *TCoolBar) EdgeInner() TEdgeStyle {
    return CoolBar_GetEdgeInner(c.instance)
}

func (c *TCoolBar) SetEdgeInner(value TEdgeStyle) {
    CoolBar_SetEdgeInner(c.instance, value)
}

func (c *TCoolBar) EdgeOuter() TEdgeStyle {
    return CoolBar_GetEdgeOuter(c.instance)
}

func (c *TCoolBar) SetEdgeOuter(value TEdgeStyle) {
    CoolBar_SetEdgeOuter(c.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TCoolBar) Enabled() bool {
    return CoolBar_GetEnabled(c.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TCoolBar) SetEnabled(value bool) {
    CoolBar_SetEnabled(c.instance, value)
}

func (c *TCoolBar) FixedSize() bool {
    return CoolBar_GetFixedSize(c.instance)
}

func (c *TCoolBar) SetFixedSize(value bool) {
    CoolBar_SetFixedSize(c.instance, value)
}

func (c *TCoolBar) FixedOrder() bool {
    return CoolBar_GetFixedOrder(c.instance)
}

func (c *TCoolBar) SetFixedOrder(value bool) {
    CoolBar_SetFixedOrder(c.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (c *TCoolBar) Font() *TFont {
    return AsFont(CoolBar_GetFont(c.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (c *TCoolBar) SetFont(value *TFont) {
    CoolBar_SetFont(c.instance, CheckPtr(value))
}

// CN: 获取图标索引列表对象。
// EN: .
func (c *TCoolBar) Images() *TImageList {
    return AsImageList(CoolBar_GetImages(c.instance))
}

// CN: 设置图标索引列表对象。
// EN: .
func (c *TCoolBar) SetImages(value IComponent) {
    CoolBar_SetImages(c.instance, CheckPtr(value))
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (c *TCoolBar) ParentColor() bool {
    return CoolBar_GetParentColor(c.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (c *TCoolBar) SetParentColor(value bool) {
    CoolBar_SetParentColor(c.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (c *TCoolBar) ParentDoubleBuffered() bool {
    return CoolBar_GetParentDoubleBuffered(c.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (c *TCoolBar) SetParentDoubleBuffered(value bool) {
    CoolBar_SetParentDoubleBuffered(c.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (c *TCoolBar) ParentFont() bool {
    return CoolBar_GetParentFont(c.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (c *TCoolBar) SetParentFont(value bool) {
    CoolBar_SetParentFont(c.instance, value)
}

func (c *TCoolBar) ParentShowHint() bool {
    return CoolBar_GetParentShowHint(c.instance)
}

func (c *TCoolBar) SetParentShowHint(value bool) {
    CoolBar_SetParentShowHint(c.instance, value)
}

func (c *TCoolBar) Bitmap() *TBitmap {
    return AsBitmap(CoolBar_GetBitmap(c.instance))
}

func (c *TCoolBar) SetBitmap(value *TBitmap) {
    CoolBar_SetBitmap(c.instance, CheckPtr(value))
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (c *TCoolBar) PopupMenu() *TPopupMenu {
    return AsPopupMenu(CoolBar_GetPopupMenu(c.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (c *TCoolBar) SetPopupMenu(value IComponent) {
    CoolBar_SetPopupMenu(c.instance, CheckPtr(value))
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (c *TCoolBar) ShowHint() bool {
    return CoolBar_GetShowHint(c.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (c *TCoolBar) SetShowHint(value bool) {
    CoolBar_SetShowHint(c.instance, value)
}

func (c *TCoolBar) ShowText() bool {
    return CoolBar_GetShowText(c.instance)
}

func (c *TCoolBar) SetShowText(value bool) {
    CoolBar_SetShowText(c.instance, value)
}

func (c *TCoolBar) Vertical() bool {
    return CoolBar_GetVertical(c.instance)
}

func (c *TCoolBar) SetVertical(value bool) {
    CoolBar_SetVertical(c.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TCoolBar) Visible() bool {
    return CoolBar_GetVisible(c.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TCoolBar) SetVisible(value bool) {
    CoolBar_SetVisible(c.instance, value)
}

// CN: 设置改变事件。
// EN: Set changed event.
func (c *TCoolBar) SetOnChange(fn TNotifyEvent) {
    CoolBar_SetOnChange(c.instance, fn)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TCoolBar) SetOnClick(fn TNotifyEvent) {
    CoolBar_SetOnClick(c.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (c *TCoolBar) SetOnContextPopup(fn TContextPopupEvent) {
    CoolBar_SetOnContextPopup(c.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (c *TCoolBar) SetOnDblClick(fn TNotifyEvent) {
    CoolBar_SetOnDblClick(c.instance, fn)
}

func (c *TCoolBar) SetOnDockDrop(fn TDockDropEvent) {
    CoolBar_SetOnDockDrop(c.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (c *TCoolBar) SetOnDragDrop(fn TDragDropEvent) {
    CoolBar_SetOnDragDrop(c.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (c *TCoolBar) SetOnDragOver(fn TDragOverEvent) {
    CoolBar_SetOnDragOver(c.instance, fn)
}

// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (c *TCoolBar) SetOnEndDock(fn TEndDragEvent) {
    CoolBar_SetOnEndDock(c.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (c *TCoolBar) SetOnEndDrag(fn TEndDragEvent) {
    CoolBar_SetOnEndDrag(c.instance, fn)
}

func (c *TCoolBar) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    CoolBar_SetOnGetSiteInfo(c.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (c *TCoolBar) SetOnMouseDown(fn TMouseEvent) {
    CoolBar_SetOnMouseDown(c.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (c *TCoolBar) SetOnMouseEnter(fn TNotifyEvent) {
    CoolBar_SetOnMouseEnter(c.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (c *TCoolBar) SetOnMouseLeave(fn TNotifyEvent) {
    CoolBar_SetOnMouseLeave(c.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (c *TCoolBar) SetOnMouseMove(fn TMouseMoveEvent) {
    CoolBar_SetOnMouseMove(c.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (c *TCoolBar) SetOnMouseUp(fn TMouseEvent) {
    CoolBar_SetOnMouseUp(c.instance, fn)
}

// CN: 设置大小被改变事件。
// EN: .
func (c *TCoolBar) SetOnResize(fn TNotifyEvent) {
    CoolBar_SetOnResize(c.instance, fn)
}

// CN: 设置启动停靠。
// EN: .
func (c *TCoolBar) SetOnStartDock(fn TStartDockEvent) {
    CoolBar_SetOnStartDock(c.instance, fn)
}

func (c *TCoolBar) SetOnUnDock(fn TUnDockEvent) {
    CoolBar_SetOnUnDock(c.instance, fn)
}

// CN: 获取依靠客户端总数。
// EN: .
func (c *TCoolBar) DockClientCount() int32 {
    return CoolBar_GetDockClientCount(c.instance)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (c *TCoolBar) MouseInClient() bool {
    return CoolBar_GetMouseInClient(c.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (c *TCoolBar) VisibleDockClientCount() int32 {
    return CoolBar_GetVisibleDockClientCount(c.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (c *TCoolBar) Brush() *TBrush {
    return AsBrush(CoolBar_GetBrush(c.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (c *TCoolBar) ControlCount() int32 {
    return CoolBar_GetControlCount(c.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (c *TCoolBar) Handle() HWND {
    return CoolBar_GetHandle(c.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (c *TCoolBar) ParentWindow() HWND {
    return CoolBar_GetParentWindow(c.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (c *TCoolBar) SetParentWindow(value HWND) {
    CoolBar_SetParentWindow(c.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (c *TCoolBar) TabOrder() TTabOrder {
    return CoolBar_GetTabOrder(c.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (c *TCoolBar) SetTabOrder(value TTabOrder) {
    CoolBar_SetTabOrder(c.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (c *TCoolBar) TabStop() bool {
    return CoolBar_GetTabStop(c.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (c *TCoolBar) SetTabStop(value bool) {
    CoolBar_SetTabStop(c.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (c *TCoolBar) UseDockManager() bool {
    return CoolBar_GetUseDockManager(c.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (c *TCoolBar) SetUseDockManager(value bool) {
    CoolBar_SetUseDockManager(c.instance, value)
}

func (c *TCoolBar) Action() *TAction {
    return AsAction(CoolBar_GetAction(c.instance))
}

func (c *TCoolBar) SetAction(value IComponent) {
    CoolBar_SetAction(c.instance, CheckPtr(value))
}

func (c *TCoolBar) BiDiMode() TBiDiMode {
    return CoolBar_GetBiDiMode(c.instance)
}

func (c *TCoolBar) SetBiDiMode(value TBiDiMode) {
    CoolBar_SetBiDiMode(c.instance, value)
}

func (c *TCoolBar) BoundsRect() TRect {
    return CoolBar_GetBoundsRect(c.instance)
}

func (c *TCoolBar) SetBoundsRect(value TRect) {
    CoolBar_SetBoundsRect(c.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (c *TCoolBar) ClientHeight() int32 {
    return CoolBar_GetClientHeight(c.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (c *TCoolBar) SetClientHeight(value int32) {
    CoolBar_SetClientHeight(c.instance, value)
}

func (c *TCoolBar) ClientOrigin() TPoint {
    return CoolBar_GetClientOrigin(c.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (c *TCoolBar) ClientRect() TRect {
    return CoolBar_GetClientRect(c.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (c *TCoolBar) ClientWidth() int32 {
    return CoolBar_GetClientWidth(c.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (c *TCoolBar) SetClientWidth(value int32) {
    CoolBar_SetClientWidth(c.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (c *TCoolBar) ControlState() TControlState {
    return CoolBar_GetControlState(c.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (c *TCoolBar) SetControlState(value TControlState) {
    CoolBar_SetControlState(c.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (c *TCoolBar) ControlStyle() TControlStyle {
    return CoolBar_GetControlStyle(c.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (c *TCoolBar) SetControlStyle(value TControlStyle) {
    CoolBar_SetControlStyle(c.instance, value)
}

func (c *TCoolBar) Floating() bool {
    return CoolBar_GetFloating(c.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TCoolBar) Parent() *TWinControl {
    return AsWinControl(CoolBar_GetParent(c.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TCoolBar) SetParent(value IWinControl) {
    CoolBar_SetParent(c.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (c *TCoolBar) Left() int32 {
    return CoolBar_GetLeft(c.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (c *TCoolBar) SetLeft(value int32) {
    CoolBar_SetLeft(c.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (c *TCoolBar) Top() int32 {
    return CoolBar_GetTop(c.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (c *TCoolBar) SetTop(value int32) {
    CoolBar_SetTop(c.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (c *TCoolBar) Width() int32 {
    return CoolBar_GetWidth(c.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (c *TCoolBar) SetWidth(value int32) {
    CoolBar_SetWidth(c.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (c *TCoolBar) Height() int32 {
    return CoolBar_GetHeight(c.instance)
}

// CN: 设置高度。
// EN: Set height.
func (c *TCoolBar) SetHeight(value int32) {
    CoolBar_SetHeight(c.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TCoolBar) Cursor() TCursor {
    return CoolBar_GetCursor(c.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TCoolBar) SetCursor(value TCursor) {
    CoolBar_SetCursor(c.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (c *TCoolBar) Hint() string {
    return CoolBar_GetHint(c.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (c *TCoolBar) SetHint(value string) {
    CoolBar_SetHint(c.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TCoolBar) ComponentCount() int32 {
    return CoolBar_GetComponentCount(c.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (c *TCoolBar) ComponentIndex() int32 {
    return CoolBar_GetComponentIndex(c.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (c *TCoolBar) SetComponentIndex(value int32) {
    CoolBar_SetComponentIndex(c.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TCoolBar) Owner() *TComponent {
    return AsComponent(CoolBar_GetOwner(c.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (c *TCoolBar) Name() string {
    return CoolBar_GetName(c.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (c *TCoolBar) SetName(value string) {
    CoolBar_SetName(c.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TCoolBar) Tag() int {
    return CoolBar_GetTag(c.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TCoolBar) SetTag(value int) {
    CoolBar_SetTag(c.instance, value)
}

func (c *TCoolBar) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(CoolBar_GetAnchorSideLeft(c.instance))
}

func (c *TCoolBar) SetAnchorSideLeft(value *TAnchorSide) {
    CoolBar_SetAnchorSideLeft(c.instance, CheckPtr(value))
}

func (c *TCoolBar) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(CoolBar_GetAnchorSideTop(c.instance))
}

func (c *TCoolBar) SetAnchorSideTop(value *TAnchorSide) {
    CoolBar_SetAnchorSideTop(c.instance, CheckPtr(value))
}

func (c *TCoolBar) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(CoolBar_GetAnchorSideRight(c.instance))
}

func (c *TCoolBar) SetAnchorSideRight(value *TAnchorSide) {
    CoolBar_SetAnchorSideRight(c.instance, CheckPtr(value))
}

func (c *TCoolBar) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(CoolBar_GetAnchorSideBottom(c.instance))
}

func (c *TCoolBar) SetAnchorSideBottom(value *TAnchorSide) {
    CoolBar_SetAnchorSideBottom(c.instance, CheckPtr(value))
}

func (c *TCoolBar) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(CoolBar_GetChildSizing(c.instance))
}

func (c *TCoolBar) SetChildSizing(value *TControlChildSizing) {
    CoolBar_SetChildSizing(c.instance, CheckPtr(value))
}

func (c *TCoolBar) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(CoolBar_GetBorderSpacing(c.instance))
}

func (c *TCoolBar) SetBorderSpacing(value *TControlBorderSpacing) {
    CoolBar_SetBorderSpacing(c.instance, CheckPtr(value))
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (c *TCoolBar) DockClients(Index int32) *TControl {
    return AsControl(CoolBar_GetDockClients(c.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (c *TCoolBar) Controls(Index int32) *TControl {
    return AsControl(CoolBar_GetControls(c.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TCoolBar) Components(AIndex int32) *TComponent {
    return AsComponent(CoolBar_GetComponents(c.instance, AIndex))
}

func (c *TCoolBar) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(CoolBar_GetAnchorSide(c.instance, AKind))
}

