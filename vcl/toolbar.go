
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

type TToolBar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewToolBar(owner IComponent) *TToolBar {
    t := new(TToolBar)
    t.instance = ToolBar_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsToolBar(obj interface{}) *TToolBar {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TToolBar{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsToolBar.
func ToolBarFromInst(inst uintptr) *TToolBar {
    return AsToolBar(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsToolBar.
func ToolBarFromObj(obj IObject) *TToolBar {
    return AsToolBar(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsToolBar.
func ToolBarFromUnsafePointer(ptr unsafe.Pointer) *TToolBar {
    return AsToolBar(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (t *TToolBar) Free() {
    if t.instance != 0 {
        ToolBar_Free(t.instance)
        t.instance, t.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TToolBar) Instance() uintptr {
    return t.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TToolBar) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TToolBar) IsValid() bool {
    return t.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (t *TToolBar) Is() TIs {
    return TIs(t.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (t *TToolBar) As() TAs {
//    return TAs(t.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TToolBarClass() TClass {
    return ToolBar_StaticClassType()
}

func (t *TToolBar) FlipChildren(AllLevels bool) {
    ToolBar_FlipChildren(t.instance, AllLevels)
}

// CN: 是否可以获得焦点。
// EN: .
func (t *TToolBar) CanFocus() bool {
    return ToolBar_CanFocus(t.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (t *TToolBar) ContainsControl(Control IControl) bool {
    return ToolBar_ContainsControl(t.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (t *TToolBar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(ToolBar_ControlAtPos(t.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (t *TToolBar) DisableAlign() {
    ToolBar_DisableAlign(t.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (t *TToolBar) EnableAlign() {
    ToolBar_EnableAlign(t.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (t *TToolBar) FindChildControl(ControlName string) *TControl {
    return AsControl(ToolBar_FindChildControl(t.instance, ControlName))
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (t *TToolBar) Focused() bool {
    return ToolBar_Focused(t.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (t *TToolBar) HandleAllocated() bool {
    return ToolBar_HandleAllocated(t.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (t *TToolBar) InsertControl(AControl IControl) {
    ToolBar_InsertControl(t.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (t *TToolBar) Invalidate() {
    ToolBar_Invalidate(t.instance)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (t *TToolBar) RemoveControl(AControl IControl) {
    ToolBar_RemoveControl(t.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (t *TToolBar) Realign() {
    ToolBar_Realign(t.instance)
}

// CN: 重绘。
// EN: Repaint.
func (t *TToolBar) Repaint() {
    ToolBar_Repaint(t.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (t *TToolBar) ScaleBy(M int32, D int32) {
    ToolBar_ScaleBy(t.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (t *TToolBar) ScrollBy(DeltaX int32, DeltaY int32) {
    ToolBar_ScrollBy(t.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (t *TToolBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ToolBar_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (t *TToolBar) SetFocus() {
    ToolBar_SetFocus(t.instance)
}

// CN: 控件更新。
// EN: Update.
func (t *TToolBar) Update() {
    ToolBar_Update(t.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (t *TToolBar) BringToFront() {
    ToolBar_BringToFront(t.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (t *TToolBar) ClientToScreen(Point TPoint) TPoint {
    return ToolBar_ClientToScreen(t.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (t *TToolBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ToolBar_ClientToParent(t.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (t *TToolBar) Dragging() bool {
    return ToolBar_Dragging(t.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (t *TToolBar) HasParent() bool {
    return ToolBar_HasParent(t.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (t *TToolBar) Hide() {
    ToolBar_Hide(t.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (t *TToolBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ToolBar_Perform(t.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (t *TToolBar) Refresh() {
    ToolBar_Refresh(t.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (t *TToolBar) ScreenToClient(Point TPoint) TPoint {
    return ToolBar_ScreenToClient(t.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (t *TToolBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ToolBar_ParentToClient(t.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (t *TToolBar) SendToBack() {
    ToolBar_SendToBack(t.instance)
}

// CN: 显示控件。
// EN: Show control.
func (t *TToolBar) Show() {
    ToolBar_Show(t.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (t *TToolBar) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return ToolBar_GetTextBuf(t.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (t *TToolBar) GetTextLen() int32 {
    return ToolBar_GetTextLen(t.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (t *TToolBar) SetTextBuf(Buffer string) {
    ToolBar_SetTextBuf(t.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (t *TToolBar) FindComponent(AName string) *TComponent {
    return AsComponent(ToolBar_FindComponent(t.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TToolBar) GetNamePath() string {
    return ToolBar_GetNamePath(t.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TToolBar) Assign(Source IObject) {
    ToolBar_Assign(t.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TToolBar) ClassType() TClass {
    return ToolBar_ClassType(t.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TToolBar) ClassName() string {
    return ToolBar_ClassName(t.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TToolBar) InstanceSize() int32 {
    return ToolBar_InstanceSize(t.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TToolBar) InheritsFrom(AClass TClass) bool {
    return ToolBar_InheritsFrom(t.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TToolBar) Equals(Obj IObject) bool {
    return ToolBar_Equals(t.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TToolBar) GetHashCode() int32 {
    return ToolBar_GetHashCode(t.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (t *TToolBar) ToString() string {
    return ToolBar_ToString(t.instance)
}

func (t *TToolBar) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ToolBar_AnchorToNeighbour(t.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (t *TToolBar) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ToolBar_AnchorParallel(t.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (t *TToolBar) AnchorHorizontalCenterTo(ASibling IControl) {
    ToolBar_AnchorHorizontalCenterTo(t.instance, CheckPtr(ASibling))
}

func (t *TToolBar) AnchorVerticalCenterTo(ASibling IControl) {
    ToolBar_AnchorVerticalCenterTo(t.instance, CheckPtr(ASibling))
}

func (t *TToolBar) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    ToolBar_AnchorAsAlign(t.instance, ATheAlign , ASpace)
}

func (t *TToolBar) AnchorClient(ASpace int32) {
    ToolBar_AnchorClient(t.instance, ASpace)
}

func (t *TToolBar) ButtonCount() int32 {
    return ToolBar_GetButtonCount(t.instance)
}

// CN: 获取画布。
// EN: .
func (t *TToolBar) Canvas() *TCanvas {
    return AsCanvas(ToolBar_GetCanvas(t.instance))
}

func (t *TToolBar) RowCount() int32 {
    return ToolBar_GetRowCount(t.instance)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (t *TToolBar) Align() TAlign {
    return ToolBar_GetAlign(t.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (t *TToolBar) SetAlign(value TAlign) {
    ToolBar_SetAlign(t.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (t *TToolBar) Anchors() TAnchors {
    return ToolBar_GetAnchors(t.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (t *TToolBar) SetAnchors(value TAnchors) {
    ToolBar_SetAnchors(t.instance, value)
}

// CN: 获取自动调整大小。
// EN: .
func (t *TToolBar) AutoSize() bool {
    return ToolBar_GetAutoSize(t.instance)
}

// CN: 设置自动调整大小。
// EN: .
func (t *TToolBar) SetAutoSize(value bool) {
    ToolBar_SetAutoSize(t.instance, value)
}

// CN: 获取边框的宽度。
// EN: .
func (t *TToolBar) BorderWidth() int32 {
    return ToolBar_GetBorderWidth(t.instance)
}

// CN: 设置边框的宽度。
// EN: .
func (t *TToolBar) SetBorderWidth(value int32) {
    ToolBar_SetBorderWidth(t.instance, value)
}

func (t *TToolBar) ButtonHeight() int32 {
    return ToolBar_GetButtonHeight(t.instance)
}

func (t *TToolBar) SetButtonHeight(value int32) {
    ToolBar_SetButtonHeight(t.instance, value)
}

func (t *TToolBar) ButtonWidth() int32 {
    return ToolBar_GetButtonWidth(t.instance)
}

func (t *TToolBar) SetButtonWidth(value int32) {
    ToolBar_SetButtonWidth(t.instance, value)
}

// CN: 获取控件标题。
// EN: Get the control title.
func (t *TToolBar) Caption() string {
    return ToolBar_GetCaption(t.instance)
}

// CN: 设置控件标题。
// EN: Set the control title.
func (t *TToolBar) SetCaption(value string) {
    ToolBar_SetCaption(t.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (t *TToolBar) Color() TColor {
    return ToolBar_GetColor(t.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (t *TToolBar) SetColor(value TColor) {
    ToolBar_SetColor(t.instance, value)
}

func (t *TToolBar) Constraints() *TSizeConstraints {
    return AsSizeConstraints(ToolBar_GetConstraints(t.instance))
}

func (t *TToolBar) SetConstraints(value *TSizeConstraints) {
    ToolBar_SetConstraints(t.instance, CheckPtr(value))
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (t *TToolBar) DoubleBuffered() bool {
    return ToolBar_GetDoubleBuffered(t.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (t *TToolBar) SetDoubleBuffered(value bool) {
    ToolBar_SetDoubleBuffered(t.instance, value)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (t *TToolBar) DockSite() bool {
    return ToolBar_GetDockSite(t.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (t *TToolBar) SetDockSite(value bool) {
    ToolBar_SetDockSite(t.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (t *TToolBar) DragCursor() TCursor {
    return ToolBar_GetDragCursor(t.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (t *TToolBar) SetDragCursor(value TCursor) {
    ToolBar_SetDragCursor(t.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (t *TToolBar) DragKind() TDragKind {
    return ToolBar_GetDragKind(t.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (t *TToolBar) SetDragKind(value TDragKind) {
    ToolBar_SetDragKind(t.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (t *TToolBar) DragMode() TDragMode {
    return ToolBar_GetDragMode(t.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (t *TToolBar) SetDragMode(value TDragMode) {
    ToolBar_SetDragMode(t.instance, value)
}

func (t *TToolBar) EdgeBorders() TEdgeBorders {
    return ToolBar_GetEdgeBorders(t.instance)
}

func (t *TToolBar) SetEdgeBorders(value TEdgeBorders) {
    ToolBar_SetEdgeBorders(t.instance, value)
}

func (t *TToolBar) EdgeInner() TEdgeStyle {
    return ToolBar_GetEdgeInner(t.instance)
}

func (t *TToolBar) SetEdgeInner(value TEdgeStyle) {
    ToolBar_SetEdgeInner(t.instance, value)
}

func (t *TToolBar) EdgeOuter() TEdgeStyle {
    return ToolBar_GetEdgeOuter(t.instance)
}

func (t *TToolBar) SetEdgeOuter(value TEdgeStyle) {
    ToolBar_SetEdgeOuter(t.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (t *TToolBar) Enabled() bool {
    return ToolBar_GetEnabled(t.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (t *TToolBar) SetEnabled(value bool) {
    ToolBar_SetEnabled(t.instance, value)
}

// CN: 获取平面样式。
// EN: .
func (t *TToolBar) Flat() bool {
    return ToolBar_GetFlat(t.instance)
}

// CN: 设置平面样式。
// EN: .
func (t *TToolBar) SetFlat(value bool) {
    ToolBar_SetFlat(t.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (t *TToolBar) Font() *TFont {
    return AsFont(ToolBar_GetFont(t.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (t *TToolBar) SetFont(value *TFont) {
    ToolBar_SetFont(t.instance, CheckPtr(value))
}

// CN: 获取高度。
// EN: Get height.
func (t *TToolBar) Height() int32 {
    return ToolBar_GetHeight(t.instance)
}

// CN: 设置高度。
// EN: Set height.
func (t *TToolBar) SetHeight(value int32) {
    ToolBar_SetHeight(t.instance, value)
}

func (t *TToolBar) HotImages() *TImageList {
    return AsImageList(ToolBar_GetHotImages(t.instance))
}

func (t *TToolBar) SetHotImages(value IComponent) {
    ToolBar_SetHotImages(t.instance, CheckPtr(value))
}

// CN: 获取图标索引列表对象。
// EN: .
func (t *TToolBar) Images() *TImageList {
    return AsImageList(ToolBar_GetImages(t.instance))
}

// CN: 设置图标索引列表对象。
// EN: .
func (t *TToolBar) SetImages(value IComponent) {
    ToolBar_SetImages(t.instance, CheckPtr(value))
}

func (t *TToolBar) Indent() int32 {
    return ToolBar_GetIndent(t.instance)
}

func (t *TToolBar) SetIndent(value int32) {
    ToolBar_SetIndent(t.instance, value)
}

func (t *TToolBar) List() bool {
    return ToolBar_GetList(t.instance)
}

func (t *TToolBar) SetList(value bool) {
    ToolBar_SetList(t.instance, value)
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (t *TToolBar) ParentColor() bool {
    return ToolBar_GetParentColor(t.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (t *TToolBar) SetParentColor(value bool) {
    ToolBar_SetParentColor(t.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (t *TToolBar) ParentDoubleBuffered() bool {
    return ToolBar_GetParentDoubleBuffered(t.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (t *TToolBar) SetParentDoubleBuffered(value bool) {
    ToolBar_SetParentDoubleBuffered(t.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (t *TToolBar) ParentFont() bool {
    return ToolBar_GetParentFont(t.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (t *TToolBar) SetParentFont(value bool) {
    ToolBar_SetParentFont(t.instance, value)
}

func (t *TToolBar) ParentShowHint() bool {
    return ToolBar_GetParentShowHint(t.instance)
}

func (t *TToolBar) SetParentShowHint(value bool) {
    ToolBar_SetParentShowHint(t.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (t *TToolBar) PopupMenu() *TPopupMenu {
    return AsPopupMenu(ToolBar_GetPopupMenu(t.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (t *TToolBar) SetPopupMenu(value IComponent) {
    ToolBar_SetPopupMenu(t.instance, CheckPtr(value))
}

// CN: 获取显示标题。
// EN: .
func (t *TToolBar) ShowCaptions() bool {
    return ToolBar_GetShowCaptions(t.instance)
}

// CN: 设置显示标题。
// EN: .
func (t *TToolBar) SetShowCaptions(value bool) {
    ToolBar_SetShowCaptions(t.instance, value)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (t *TToolBar) ShowHint() bool {
    return ToolBar_GetShowHint(t.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (t *TToolBar) SetShowHint(value bool) {
    ToolBar_SetShowHint(t.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (t *TToolBar) TabOrder() TTabOrder {
    return ToolBar_GetTabOrder(t.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (t *TToolBar) SetTabOrder(value TTabOrder) {
    ToolBar_SetTabOrder(t.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (t *TToolBar) TabStop() bool {
    return ToolBar_GetTabStop(t.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (t *TToolBar) SetTabStop(value bool) {
    ToolBar_SetTabStop(t.instance, value)
}

// CN: 获取透明。
// EN: Get transparent.
func (t *TToolBar) Transparent() bool {
    return ToolBar_GetTransparent(t.instance)
}

// CN: 设置透明。
// EN: Set transparent.
func (t *TToolBar) SetTransparent(value bool) {
    ToolBar_SetTransparent(t.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (t *TToolBar) Visible() bool {
    return ToolBar_GetVisible(t.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (t *TToolBar) SetVisible(value bool) {
    ToolBar_SetVisible(t.instance, value)
}

func (t *TToolBar) Wrapable() bool {
    return ToolBar_GetWrapable(t.instance)
}

func (t *TToolBar) SetWrapable(value bool) {
    ToolBar_SetWrapable(t.instance, value)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (t *TToolBar) SetOnClick(fn TNotifyEvent) {
    ToolBar_SetOnClick(t.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (t *TToolBar) SetOnContextPopup(fn TContextPopupEvent) {
    ToolBar_SetOnContextPopup(t.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (t *TToolBar) SetOnDblClick(fn TNotifyEvent) {
    ToolBar_SetOnDblClick(t.instance, fn)
}

func (t *TToolBar) SetOnDockDrop(fn TDockDropEvent) {
    ToolBar_SetOnDockDrop(t.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (t *TToolBar) SetOnDragDrop(fn TDragDropEvent) {
    ToolBar_SetOnDragDrop(t.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (t *TToolBar) SetOnDragOver(fn TDragOverEvent) {
    ToolBar_SetOnDragOver(t.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (t *TToolBar) SetOnEndDrag(fn TEndDragEvent) {
    ToolBar_SetOnEndDrag(t.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (t *TToolBar) SetOnEnter(fn TNotifyEvent) {
    ToolBar_SetOnEnter(t.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (t *TToolBar) SetOnExit(fn TNotifyEvent) {
    ToolBar_SetOnExit(t.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (t *TToolBar) SetOnMouseDown(fn TMouseEvent) {
    ToolBar_SetOnMouseDown(t.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (t *TToolBar) SetOnMouseEnter(fn TNotifyEvent) {
    ToolBar_SetOnMouseEnter(t.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (t *TToolBar) SetOnMouseLeave(fn TNotifyEvent) {
    ToolBar_SetOnMouseLeave(t.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (t *TToolBar) SetOnMouseMove(fn TMouseMoveEvent) {
    ToolBar_SetOnMouseMove(t.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (t *TToolBar) SetOnMouseUp(fn TMouseEvent) {
    ToolBar_SetOnMouseUp(t.instance, fn)
}

// CN: 设置大小被改变事件。
// EN: .
func (t *TToolBar) SetOnResize(fn TNotifyEvent) {
    ToolBar_SetOnResize(t.instance, fn)
}

func (t *TToolBar) SetOnUnDock(fn TUnDockEvent) {
    ToolBar_SetOnUnDock(t.instance, fn)
}

// CN: 获取依靠客户端总数。
// EN: .
func (t *TToolBar) DockClientCount() int32 {
    return ToolBar_GetDockClientCount(t.instance)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (t *TToolBar) MouseInClient() bool {
    return ToolBar_GetMouseInClient(t.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (t *TToolBar) VisibleDockClientCount() int32 {
    return ToolBar_GetVisibleDockClientCount(t.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (t *TToolBar) Brush() *TBrush {
    return AsBrush(ToolBar_GetBrush(t.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (t *TToolBar) ControlCount() int32 {
    return ToolBar_GetControlCount(t.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (t *TToolBar) Handle() HWND {
    return ToolBar_GetHandle(t.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (t *TToolBar) ParentWindow() HWND {
    return ToolBar_GetParentWindow(t.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (t *TToolBar) SetParentWindow(value HWND) {
    ToolBar_SetParentWindow(t.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (t *TToolBar) UseDockManager() bool {
    return ToolBar_GetUseDockManager(t.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (t *TToolBar) SetUseDockManager(value bool) {
    ToolBar_SetUseDockManager(t.instance, value)
}

func (t *TToolBar) Action() *TAction {
    return AsAction(ToolBar_GetAction(t.instance))
}

func (t *TToolBar) SetAction(value IComponent) {
    ToolBar_SetAction(t.instance, CheckPtr(value))
}

func (t *TToolBar) BiDiMode() TBiDiMode {
    return ToolBar_GetBiDiMode(t.instance)
}

func (t *TToolBar) SetBiDiMode(value TBiDiMode) {
    ToolBar_SetBiDiMode(t.instance, value)
}

func (t *TToolBar) BoundsRect() TRect {
    return ToolBar_GetBoundsRect(t.instance)
}

func (t *TToolBar) SetBoundsRect(value TRect) {
    ToolBar_SetBoundsRect(t.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (t *TToolBar) ClientHeight() int32 {
    return ToolBar_GetClientHeight(t.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (t *TToolBar) SetClientHeight(value int32) {
    ToolBar_SetClientHeight(t.instance, value)
}

func (t *TToolBar) ClientOrigin() TPoint {
    return ToolBar_GetClientOrigin(t.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (t *TToolBar) ClientRect() TRect {
    return ToolBar_GetClientRect(t.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (t *TToolBar) ClientWidth() int32 {
    return ToolBar_GetClientWidth(t.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (t *TToolBar) SetClientWidth(value int32) {
    ToolBar_SetClientWidth(t.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (t *TToolBar) ControlState() TControlState {
    return ToolBar_GetControlState(t.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (t *TToolBar) SetControlState(value TControlState) {
    ToolBar_SetControlState(t.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (t *TToolBar) ControlStyle() TControlStyle {
    return ToolBar_GetControlStyle(t.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (t *TToolBar) SetControlStyle(value TControlStyle) {
    ToolBar_SetControlStyle(t.instance, value)
}

func (t *TToolBar) Floating() bool {
    return ToolBar_GetFloating(t.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (t *TToolBar) Parent() *TWinControl {
    return AsWinControl(ToolBar_GetParent(t.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (t *TToolBar) SetParent(value IWinControl) {
    ToolBar_SetParent(t.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (t *TToolBar) Left() int32 {
    return ToolBar_GetLeft(t.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (t *TToolBar) SetLeft(value int32) {
    ToolBar_SetLeft(t.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (t *TToolBar) Top() int32 {
    return ToolBar_GetTop(t.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (t *TToolBar) SetTop(value int32) {
    ToolBar_SetTop(t.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (t *TToolBar) Width() int32 {
    return ToolBar_GetWidth(t.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (t *TToolBar) SetWidth(value int32) {
    ToolBar_SetWidth(t.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (t *TToolBar) Cursor() TCursor {
    return ToolBar_GetCursor(t.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (t *TToolBar) SetCursor(value TCursor) {
    ToolBar_SetCursor(t.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (t *TToolBar) Hint() string {
    return ToolBar_GetHint(t.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (t *TToolBar) SetHint(value string) {
    ToolBar_SetHint(t.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TToolBar) ComponentCount() int32 {
    return ToolBar_GetComponentCount(t.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (t *TToolBar) ComponentIndex() int32 {
    return ToolBar_GetComponentIndex(t.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (t *TToolBar) SetComponentIndex(value int32) {
    ToolBar_SetComponentIndex(t.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TToolBar) Owner() *TComponent {
    return AsComponent(ToolBar_GetOwner(t.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (t *TToolBar) Name() string {
    return ToolBar_GetName(t.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (t *TToolBar) SetName(value string) {
    ToolBar_SetName(t.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TToolBar) Tag() int {
    return ToolBar_GetTag(t.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TToolBar) SetTag(value int) {
    ToolBar_SetTag(t.instance, value)
}

func (t *TToolBar) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(ToolBar_GetAnchorSideLeft(t.instance))
}

func (t *TToolBar) SetAnchorSideLeft(value *TAnchorSide) {
    ToolBar_SetAnchorSideLeft(t.instance, CheckPtr(value))
}

func (t *TToolBar) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(ToolBar_GetAnchorSideTop(t.instance))
}

func (t *TToolBar) SetAnchorSideTop(value *TAnchorSide) {
    ToolBar_SetAnchorSideTop(t.instance, CheckPtr(value))
}

func (t *TToolBar) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(ToolBar_GetAnchorSideRight(t.instance))
}

func (t *TToolBar) SetAnchorSideRight(value *TAnchorSide) {
    ToolBar_SetAnchorSideRight(t.instance, CheckPtr(value))
}

func (t *TToolBar) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(ToolBar_GetAnchorSideBottom(t.instance))
}

func (t *TToolBar) SetAnchorSideBottom(value *TAnchorSide) {
    ToolBar_SetAnchorSideBottom(t.instance, CheckPtr(value))
}

func (t *TToolBar) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(ToolBar_GetChildSizing(t.instance))
}

func (t *TToolBar) SetChildSizing(value *TControlChildSizing) {
    ToolBar_SetChildSizing(t.instance, CheckPtr(value))
}

func (t *TToolBar) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(ToolBar_GetBorderSpacing(t.instance))
}

func (t *TToolBar) SetBorderSpacing(value *TControlBorderSpacing) {
    ToolBar_SetBorderSpacing(t.instance, CheckPtr(value))
}

func (t *TToolBar) Buttons(Index int32) *TToolButton {
    return AsToolButton(ToolBar_GetButtons(t.instance, Index))
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (t *TToolBar) DockClients(Index int32) *TControl {
    return AsControl(ToolBar_GetDockClients(t.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (t *TToolBar) Controls(Index int32) *TControl {
    return AsControl(ToolBar_GetControls(t.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TToolBar) Components(AIndex int32) *TComponent {
    return AsComponent(ToolBar_GetComponents(t.instance, AIndex))
}

func (t *TToolBar) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(ToolBar_GetAnchorSide(t.instance, AKind))
}

