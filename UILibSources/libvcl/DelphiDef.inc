
// ö�ٶ���4�ֽڶ���
{$Z4+}

type
  // �ض�����
  TUnixDateTime = type LongInt; // int64��ΪLongInt��ֻ��ȷ����

function ToUnixTime(ADateTime: TDateTime): TUnixDateTime; inline;
begin
  Result := DateTimeToUnix(ADateTime, False);
end;

function UnixToTime(ADateTime: TUnixDateTime): TDateTime; inline;
begin
  Result := UnixToDateTime(ADateTime, False);
end;

// SysLocale  14byte
type  // ���¶��壬Ŀ��Ϊ����go�еĽṹ
  TSysLocaleRedim = record
    DefaultLCID: TLocaleID;
    PriLangID: Integer;
    SubLangID: Integer;
    FarEast: Boolean;
    MiddleEast: Boolean;
  end;
// SysLocale
procedure DSysLocale(var AInfo: TSysLocaleRedim); stdcall;
begin
  AInfo.DefaultLCID := SysLocale.DefaultLCID;
  AInfo.PriLangID := SysLocale.PriLangID;
  AInfo.SubLangID := SysLocale.SubLangID;
  AInfo.FarEast := SysLocale.FarEast;
  AInfo.MiddleEast := SysLocale.MiddleEast;
end;


// TApplication

function Application_Instance: TApplication; stdcall;
begin
  Result := Application;
end;

function Application_CreateForm(App: TApplication; AInitScale: LongBool): TGoForm; stdcall;
begin
  LockInitScale;
  try
    SetInitScale(AInitScale);
    App.CreateForm(TGoForm, Result);
    SetInitScale(False);
  finally
    UnLockInitScale;
  end;
end;

procedure RemoveAllForms;
var
  C: TComponent;
  I: Integer;
begin
  for I := Application.ComponentCount - 1 downto 0 do
  begin
    C := Application.Components[I];
    if C is TCustomForm then
      C.Free;
  end;
end;

procedure Application_Run(AObj: TApplication); stdcall;
begin
  AObj.Run;
  // ��Run�������ֶ��ͷ�Application�е�Form����Ҫ�ǿ����ڷ�Delphi���л�������ɵ�һЩ��Դ��������
  RemoveAllForms;
end;

procedure Application_Initialize(AObj: TApplication); stdcall;
begin
  if AObj.Handle = 0 then
    AObj.CreateHandle;
  AObj.Initialize;
end;

// TForm

procedure Form_EnabledMaximize(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  if AValue then
  begin
    if not(biMaximize in AObj.BorderIcons) then
      AObj.BorderIcons := AObj.BorderIcons + [biMaximize]
  end else
  begin
    if biMaximize in AObj.BorderIcons then
      AObj.BorderIcons := AObj.BorderIcons - [biMaximize]
  end;
end;

procedure Form_EnabledMinimize(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  if AValue then
  begin
    if not(biMinimize in AObj.BorderIcons) then
      AObj.BorderIcons := AObj.BorderIcons + [biMinimize]
  end else
  begin
    if biMinimize in AObj.BorderIcons then
      AObj.BorderIcons := AObj.BorderIcons - [biMinimize]
  end;
end;

procedure Form_EnabledSystemMenu(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  if AValue then
  begin
    if not(biSystemMenu in AObj.BorderIcons) then
      AObj.BorderIcons := AObj.BorderIcons + [biSystemMenu]
  end else
  begin
    if biSystemMenu in AObj.BorderIcons then
      AObj.BorderIcons := AObj.BorderIcons - [biSystemMenu]
  end;
end;

procedure Form_SetAllowDropFiles(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.AllowDropFiles := AValue;
end;

function Form_GetAllowDropFiles(AObj: TGoForm): LongBool; stdcall;
begin
  Result := AObj.AllowDropFiles;
end;

procedure Form_SetOnDropFiles(AObj: TGoForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDropFiles := nil;
    TEventClass.Remove(AObj, geDropFiles);
    Exit;
  end;
  AObj.OnDropFiles := TEventClass.FormOnDropFiles;
  TEventClass.Add(AObj, geDropFiles, AEventId);
end;

procedure Form_SetOnDestroy(AObj: TGoForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDestroy := nil;
    TEventClass.Remove(AObj, geDestroy);
    Exit;
  end;
  AObj.OnDestroy := TEventClass.OnDestroy;
  TEventClass.Add(AObj, geDestroy, AEventId);
end;

procedure Form_SetOnConstrainedResize(AObj: TGoForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnConstrainedResize := nil;
    TEventClass.Remove(AObj, geConstrainedResize);
    Exit;
  end;
  AObj.OnConstrainedResize := TEventClass.OnConstrainedResize;
  TEventClass.Add(AObj, geConstrainedResize, AEventId);
end;

procedure Form_SetOnDeactivate(AObj: TGoForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDeactivate := nil;
    TEventClass.Remove(AObj, geDeactivate);
    Exit;
  end;
  AObj.OnDeactivate := TEventClass.OnDeactivate;
  TEventClass.Add(AObj, geDeactivate, AEventId);
end;

procedure Form_SetOnActivate(AObj: TGoForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnActivate := nil;
    TEventClass.Remove(AObj, geActivate);
    Exit;
  end;
  AObj.OnActivate := TEventClass.OnActivate;
  TEventClass.Add(AObj, geActivate, AEventId);
end;

// mouse
function Mouse_Instance: TMouse; stdcall;
begin
  Result := Mouse;
end;

// screen
function Screen_Instance: TScreen; stdcall;
begin
  Result := Screen;
end;

function DTextToShortCut(AText: PChar): TShortCut; stdcall;
begin
  Result := TextToShortCut(AText);
end;

function DShortCutToText(AVal: TShortCut): PChar; stdcall;
begin
  Result := PChar(ShortCutToText(AVal));
end;

// TClipboard
function Clipboard_Instance: TClipboard; stdcall;
begin
  Result := Clipboard;
end;

function Clipboard_SetClipboard(NewClipboard: TClipboard): TClipboard; stdcall;
begin
  Result := SetClipboard(NewClipboard);
end;

// TMemoryStream
// function Write(const Buffer; Count: Longint): Longint;
function MemoryStream_Write(AObj: TMemoryStream; Buffer: Pointer; Count: Longint): Longint; stdcall;
begin
  Result := AObj.Write(Buffer^, Count);
end;

// function Read(var Buffer; Count: Longint): Longint;
function MemoryStream_Read(AObj: TMemoryStream; Buffer: Pointer; Count: Longint): Longint; stdcall;
begin
  Result := AObj.Read(Buffer^, Count);
end;

// TCanvas
// Canvas_BrushCopy
// Canvas_CopyRect
// Canvas_Draw1
// Canvas_Draw2
// Canvas_DrawFocusRect
// Canvas_FillRect
// Canvas_FrameRect
// Canvas_StretchDraw
// Canvas_TextRect1
// Canvas_TextRect2

procedure Canvas_BrushCopy(AObj: TCanvas; var Dest: TRect; Bitmap: TBitmap;
      var Source: TRect; Color: TColor); stdcall;
begin
  AObj.BrushCopy(Dest, Bitmap, Source, Color);
end;

procedure Canvas_CopyRect(AObj: TCanvas; var Dest: TRect; Canvas: TCanvas;
  var Source: TRect); stdcall;
begin
  AObj.CopyRect(Dest, Canvas, Source);
end;

procedure Canvas_Draw1(AObj: TCanvas; X, Y: Integer; Graphic: TGraphic); stdcall;
begin
  AObj.Draw(X, Y, Graphic);
end;

procedure Canvas_Draw2(AObj: TCanvas; X, Y: Integer; Graphic: TGraphic; Opacity: Byte); stdcall;
begin
  AObj.Draw(X, Y, Graphic, Opacity);
end;

procedure Canvas_DrawFocusRect(AObj: TCanvas; var ARect: TRect); stdcall;
begin
  AObj.DrawFocusRect(ARect);
end;

procedure Canvas_FillRect(AObj: TCanvas; var Rect: TRect); stdcall;
begin
  AObj.FillRect(Rect);
end;

procedure Canvas_FrameRect(AObj: TCanvas; var Rect: TRect); stdcall;
begin
  AObj.FrameRect(Rect);
end;

procedure Canvas_StretchDraw(AObj: TCanvas; var Rect: TRect; Graphic: TGraphic); stdcall;
begin
  AObj.StretchDraw(Rect, Graphic);
end;

procedure Canvas_TextRect1(AObj: TCanvas; var Rect: TRect; X, Y: Integer; const Text: PChar); stdcall;
begin
  AObj.TextRect(Rect, X, Y, Text);
end;

function Canvas_TextRect2(AObj: TCanvas; var Rect: TRect; Text: PChar; var AOutStr: PChar; TextFormat: TTextFormat): Integer; stdcall;
var
  S: string;
begin
  Result := 0;
  S := Text;
  AObj.TextRect(Rect, S, TextFormat);
  if tfModifyString in TextFormat then
  begin
    AOutStr := PChar(S);
    Result := 1;
  end;
end;

// TImageList
// ImageList_Draw1
// ImageList_Draw2
// ImageList_DrawOverlay1
// ImageList_DrawOverlay2
// ImageList_GetIcon1
// ImageList_GetIcon2

procedure ImageList_Draw1(AObj: TImageList; Canvas: TCanvas; X, Y, Index: Integer; Enabled: LongBool); stdcall;
begin
  AObj.Draw(Canvas, X, Y, Index, Enabled);
end;

procedure ImageList_Draw2(AObj: TImageList; Canvas: TCanvas; X, Y, Index: Integer;
  ADrawingStyle: TDrawingStyle; AImageType: TImageType;
  Enabled: LongBool); stdcall;
begin
  AObj.Draw(Canvas, X, Y, Index, ADrawingStyle, AImageType, Enabled);
end;
  
procedure ImageList_DrawOverlay1(AObj: TImageList; Canvas: TCanvas; X, Y: Integer;
  ImageIndex: Integer; Overlay: TOverlay; Enabled: LongBool); stdcall;
begin
  AObj.DrawOverlay(Canvas, X, Y, ImageIndex, Overlay, Enabled);
end;
  
procedure ImageList_DrawOverlay2(AObj: TImageList; Canvas: TCanvas; X, Y: Integer;
  ImageIndex: Integer; Overlay: TOverlay; ADrawingStyle: TDrawingStyle;
  AImageType: TImageType; Enabled: LongBool); stdcall;
begin
  AObj.DrawOverlay(Canvas, X, Y, ImageIndex, Overlay, ADrawingStyle, AImageType, Enabled);
end;
  
procedure ImageList_GetIcon1(AObj: TImageList; Index: Integer; Image: TIcon); stdcall;
begin
  AObj.GetIcon(Index, Image);
end;

procedure ImageList_GetIcon2(AObj: TImageList; Index: Integer; Image: TIcon; ADrawingStyle: TDrawingStyle; AImageType: TImageType); stdcall;
begin
  AObj.GetIcon(Index, Image, ADrawingStyle, AImageType);
end;
	
 
// Other

procedure SetEventCallback(APtr: Pointer); stdcall;
begin
  GCallbackPtr := APtr;
end;

procedure DGetParam(Index: NativeInt; AArr: Pointer; out P: TGoParam); stdcall;
begin
  P := PGoParam(PByte(AArr) + Index * SizeOf(TGoParam))^;
end;

// ��string�����л�ȡ��Ա
function DGetStringArrOf(P: Pointer; AIndex: NativeInt): PChar; stdcall;
begin
  Result := PChar(PString(PByte(P) + AIndex * SizeOf(string))^);
end;

function DStrLen(p: PChar): NativeInt; stdcall;
begin
  Result := Length(p);
end;

procedure DMove(Src, Dest: Pointer; Len: NativeInt); stdcall;
begin
  Move(Src^, Dest^, Len);
end;

function DGetClassName(Obj: TObject): PChar; stdcall;
begin
  Result := PChar(Obj.ClassName);
end;

procedure DShowMessage(AMsg: PChar); stdcall;
begin
  ShowMessage(AMsg);
end;

function DGetMainInstance: HINST; stdcall;
begin
  Result := GetModuleHandle(nil);
end;

function DMessageDlg(const Msg: PChar; DlgType: TMsgDlgType;
  Buttons: TMsgDlgButtons; HelpCtx: Longint): Integer; stdcall;
begin
  Result := MessageDlg(Msg, DlgType, Buttons, HelpCtx);
end;

procedure DSetReportMemoryLeaksOnShutdown(Value: Boolean); stdcall;
begin
{$IFDEF MSWINDOWS}
  ReportMemoryLeaksOnShutdown := Value;
{$ENDIF MSWINDOWS}
end;


// �����߳�������
procedure DSynchronize(AId: NativeUInt); stdcall;
begin
  TThread.Synchronize(nil,
    procedure
    begin
      GCallbackPtr(AId, nil, 0);
    end);
end;

procedure DSysOpen(FileName: PChar); stdcall;
begin
  ShellExecute(0, nil, FileName, nil, nil, SW_SHOW);
end;

// ���뵼��������Ըɴ���delphi�ĵ���
function DExtractFilePath(AFileName: PChar): PChar; stdcall;
begin
  Result := PChar(ExtractFilePath(AFileName));
end;

function DFileExists(AFileName: PChar): LongBool; stdcall;
begin
  Result := FileExists(AFileName);
end;

function DSelectDirectory1(var Directory: PChar;
  Options: TSelectDirOpts; HelpCtx: Longint): LongBool; stdcall;
var
  LDir: string;
begin
  Result := SelectDirectory(LDir, Options, HelpCtx);
  if Result then
    Directory := PChar(LDir);
end;

function DSelectDirectory2(Caption, Root: PChar;
  var Directory: PChar; Options: TSelectDirExtOpts; Parent: TWinControl): LongBool;  stdcall;
var
  LDir: string;
begin
  Result := SelectDirectory(Caption, Root, LDir, Options, Parent);
  if Result then
    Directory := PChar(LDir);
end;

function DInputBox(ACaption, APrompt, ADefault: PChar): PChar; stdcall;
begin
  Result := PChar(InputBox(ACaption, APrompt, ADefault));
end;

function DInputQuery(ACaption, APrompt: PChar; Value: PChar; out AOut: PChar): LongBool; stdcall;
var
  S: string;
begin
  S := Value;
  Result := InputQuery(string(ACaption), string(APrompt), S);
  if Result then
    AOut := PChar(S);
end;

// ·���ϲ�
function PathCombine(A1, A2: string): string; inline;
begin
  if A1[Length(A1)] <> PathDelim then
    A1 := A1 + PathDelim;
  if (Length(A2) > 0) and (A2[1] = PathDelim) then
    Delete(A2, 1, 1);
  Result := A1 + A2;
end;

// ����URL��ݷ�ʽ
procedure DCreateURLShortCut(ADestPath, AShortCutName, AURL: PChar); stdcall;
var
  Ini: TIniFile;
begin
  Ini := TIniFile.Create(PathCombine(ADestPath, AShortCutName) + '.url');
  try
    Ini.WriteString('{000214A0-0000-0000-C000-000000000046}', 'Prop3', '19,2');
    Ini.WriteString('InternetShortcut', 'URL', AURL);
  finally
    Ini.Free;
  end;
end;

// ������ݷ�ʽ
function DCreateShortCut(ADestPath, AShortCutName, ASrcFileName, AIconFileName, ADescription, ACmdArgs: PChar): LongBool; stdcall;
var
  LObject : IUnknown;
  LPLink : IShellLink;
  LPPFile : IPersistFile;
  LinkFilename : WideString;
begin
  CoInitializeEx(nil, 0);
  try
    LObject := CreateComObject(CLSID_ShellLink);
    LPLink := LObject as IShellLink;
    LPPFile := LObject as IPersistFile;
    LPLink.SetPath(ASrcFileName);
    LPLink.SetWorkingDirectory(PWideChar(ExtractFilePath(ASrcFileName)));
    LPLink.SetDescription(ADescription);
    if AIconFileName = '' then
      LPLink.SetIconLocation(ASrcFileName, 0)
    else LPLink.SetIconLocation(AIconFileName, 0);
    LPLink.SetArguments(ACmdArgs);
    LinkFilename := WideString(PathCombine(ADestPath, AShortCutName) + '.lnk');
    Result := LPPFile.Save(PWideChar(LinkFilename), FALSE) = S_OK;
  finally
    CoUninitialize;
  end;
end;

// �������ã�������Ҫ�ṩ��������ʹ��
procedure DSetPropertyValue(Instance: TObject; PropName, Value: PChar); stdcall;
var
  LPropInfo: PPropInfo;
begin
  LPropInfo := GetPropInfo(Instance, PropName);
  if LPropInfo <> nil then
    SetStrProp(Instance, LPropInfo, Value);
end;

procedure DSetPropertySecValue(Instance: TObject; PropName, SecPropName, Value: PChar); stdcall;
type
  TCallIndexGetClassProc = function(AData: Pointer; AIndex: Integer): Pointer;
var
  LPropInfo: PPropInfo;
  LObj: TObject;
  LPropName: string;
  LP, LIndex: Integer;
begin
  LPropName := PropName;
  LIndex := -1;
  LP := LPropName.IndexOf('[');
  if LP > 0 then
  begin
    LIndex := StrToIntDef(LPropName.Substring(LP + 1, LPropName.IndexOf(']') - LP - 1), -1);
    LPropName := LPropName.Substring(0, LP);
  end;
  LPropInfo := GetPropInfo(Instance, LPropName);
  if LPropInfo <> nil then
  begin
    LObj := GetObjectProp(Instance, LPropInfo);
    if LObj <> nil then
    begin
      if LIndex <> -1 then
      begin
        if LPropInfo^.PropType^.Kind = tkClass then
        begin
          if (LObj is TListColumns) and (LIndex >= 0) and (LIndex < TListColumns(LObj).Count) and string(SecPropName).Equals('Caption') then
            TListColumns(LObj)[LIndex].Caption := Value;
        end;
      end else
        DSetPropertyValue(LObj, SecPropName, Value);
    end;
  end;
end;



// TStyleManager
// StyleManager_Initialize
// StyleManager_UnInitialize
// StyleManager_IsValidStyle
// StyleManager_LoadFromFile
// StyleManager_CheckSysClassName
// StyleManager_TrySetStyle
// StyleManager_SetStyle1
// StyleManager_SetStyle2
// StyleManager_TryLoadFromResource

function StyleManager_IsValidStyle(AFileName: PChar): LongBool; stdcall;
begin
  Result := TStyleManager.IsValidStyle(AFileName);
end;

function StyleManager_IsValidStyle2(const AFileName: PChar; var AName: PChar): LongBool; stdcall;
var
  LStyleInfo: TStyleInfo;
begin
  Result := TStyleManager.IsValidStyle(AFileName, LStyleInfo);
  if Result then
    AName := PChar(LStyleInfo.Name);
end;


function StyleManager_LoadFromFile(AFileName: PChar): TStyleManager.TStyleServicesHandle; stdcall;
begin
  Result := TStyleManager.LoadFromFile(AFileName);
end;

function StyleManager_CheckSysClassName(AClassName: PChar): LongBool; stdcall;
begin
  Result := TStyleManager.CheckSysClassName(AClassName);
end;

function StyleManager_TrySetStyle(AName: PChar; ShowErrorDialog: LongBool): LongBool; stdcall;
begin
  Result := TStyleManager.TrySetStyle(AName, ShowErrorDialog);
end;

procedure StyleManager_SetStyle1(Handle: TStyleManager.TStyleServicesHandle); stdcall;
begin
  TStyleManager.SetStyle(Handle);
end;

procedure StyleManager_SetStyle2(AName: PChar); stdcall;
begin
  TStyleManager.SetStyle(string(AName));
end;

function StyleManager_TryLoadFromResource(Instance: HINST; ResourceName: PChar;
   ResourceType: PChar; var Handle: TStyleManager.TStyleServicesHandle): LongBool; stdcall;
begin
  Result := TStyleManager.TryLoadFromResource(Instance, ResourceName, ResourceType, Handle);
end;

// StyleManager_ActiveStyle
// StyleManager_SystemStyle
// StyleManager_Enabled
// StyleManager_IsCustomStyleActive
// StyleManager_UnRegisterStyle
// StyleManager_RegisterStyle
// StyleManager_Style
// StyleManager_StyleDescriptor
function StyleManager_ActiveStyle: TCustomStyleServices; stdcall;
begin
  Result := TStyleManager.ActiveStyle;
end;

function StyleManager_SystemStyle: TCustomStyleServices; stdcall;
begin
  Result := TStyleManager.SystemStyle;
end;

function StyleManager_Enabled: LongBool; stdcall;
begin
  Result := TStyleManager.Enabled;
end;

function StyleManager_IsCustomStyleActive: LongBool; stdcall;
begin
  Result := TStyleManager.IsCustomStyleActive;
end;

procedure StyleManager_UnRegisterStyle(Style: TCustomStyleServices); stdcall;
begin
  TStyleManager.UnRegisterStyle(Style);
end;

procedure StyleManager_RegisterStyle(Style: TCustomStyleServices); stdcall;
begin
  TStyleManager.RegisterStyle(Style);
end;

function StyleManager_Style(name: PChar): TCustomStyleServices; stdcall;
begin
  Result := TStyleManager.Style[name];
end;

function StyleManager_StyleDescriptor(StyleName: PChar): TStyleManager.TStyleClassDescriptor; stdcall;
begin
  Result := TStyleManager.StyleDescriptor[StyleName];
end;

// StyleNames
//function StyleManager_StyleNamesPtr: Pointer; stdcall;
//begin
//  Result := Pointer(@TStyleManager.StyleNames[0]);
//end;
//
function StyleManager_StyleNamesOf(AIndex: Integer): PChar; stdcall;
var
  LArr: TArray<string>;
begin
  Result := nil;
  LArr := TStyleManager.StyleNames;
  if (Length(LArr) > 0) and (AIndex >= 0) and (AIndex < Length(LArr)) then
    Result := PChar(LArr[AIndex]);
end;

// dylib
type
  TSyscall0  = function: UInt64; stdcall;
  TSyscall1  = function(A1: Pointer): UInt64; stdcall;
  TSyscall2  = function(A1, A2: Pointer): UInt64; stdcall;
  TSyscall3  = function(A1, A2, A3: Pointer): UInt64; stdcall;
  TSyscall4  = function(A1, A2, A3, A4: Pointer): UInt64; stdcall;
  TSyscall5  = function(A1, A2, A3, A4, A5: Pointer): UInt64; stdcall;
  TSyscall6  = function(A1, A2, A3, A4, A5, A6: Pointer): UInt64; stdcall;
  TSyscall7  = function(A1, A2, A3, A4, A5, A6, A7: Pointer): UInt64; stdcall;
  TSyscall8  = function(A1, A2, A3, A4, A5, A6, A7, A8: Pointer): UInt64; stdcall;
  TSyscall9  = function(A1, A2, A3, A4, A5, A6, A7, A8, A9: Pointer): UInt64; stdcall;
  TSyscall10 = function(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10: Pointer): UInt64; stdcall;
  TSyscall11 = function(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11: Pointer): UInt64; stdcall;
  TSyscall12 = function(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12: Pointer): UInt64; stdcall;

function MySyscall(AProc:Pointer; ALen: NativeInt; A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12: Pointer): UInt64; stdcall;
begin
  Result := 0;
  if AProc = nil then Exit;
  try
    case ALen of
      0 : Result := TSyscall0(AProc)();
      1 : Result := TSyscall1(AProc) (A1);
      2 : Result := TSyscall2(AProc) (A1, A2);
      3 : Result := TSyscall3(AProc) (A1, A2, A3);
      4 : Result := TSyscall4(AProc) (A1, A2, A3, A4);
      5 : Result := TSyscall5(AProc) (A1, A2, A3, A4, A5);
      6 : Result := TSyscall6(AProc) (A1, A2, A3, A4, A5, A6);
      7 : Result := TSyscall7(AProc) (A1, A2, A3, A4, A5, A6, A7);
      8 : Result := TSyscall8(AProc) (A1, A2, A3, A4, A5, A6, A7, A8);
      9 : Result := TSyscall9(AProc) (A1, A2, A3, A4, A5, A6, A7, A8, A9);
      10: Result := TSyscall10(AProc)(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10);
      11: Result := TSyscall11(AProc)(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11);
      12: Result := TSyscall12(AProc)(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12);
    else
      Exit;
    end;
  except
    on E: Exception do
    begin
      if Assigned(Application) then
      begin
        if Assigned(Application.OnException) then
          Application.OnException(Application, E)
        else
          MessageDlg(E.Message,  mtError, [mbOK], 0);
      end else
        OutputDebugString(PChar('Syscall' + ALen.ToString + ' Error: ' + E.Message));
    end;
  end;
end;



exports
  // ������̬���õ�
  MySyscall,

  // SysLocaled
  DSysLocale,

  // TApplication
  Application_Instance,
  Application_CreateForm,
  Application_Run,
  Application_Initialize,

  // TForm
  Form_EnabledMaximize,
  Form_EnabledMinimize,
  Form_EnabledSystemMenu,
  Form_SetAllowDropFiles,
  Form_GetAllowDropFiles,
  Form_SetOnDropFiles,
  Form_SetOnDestroy,
  Form_SetOnConstrainedResize,
  Form_SetOnDeactivate,
  Form_SetOnActivate,

  // TMouse
  Mouse_Instance,
  Screen_Instance,

  // other
  SetEventCallback,
  DGetParam,
  DGetStringArrOf,
  DStrLen,
  DMove,
  DGetClassName,
  DShowMessage,
  DGetMainInstance,
  DSysOpen,
  DTextToShortCut,
  DShortCutToText,
  DMessageDlg,
  DExtractFilePath,
  DFileExists,

  DSelectDirectory1,
  DSelectDirectory2,
  DInputBox,
  DInputQuery,

  DSetReportMemoryLeaksOnShutdown,
  DSynchronize,

  // ShortCut
  DCreateURLShortCut,
  DCreateShortCut,

  // SetProperty
  DSetPropertyValue,
  DSetPropertySecValue,

  // TMemoryStream
  MemoryStream_Write,
  MemoryStream_Read,

  // TClipboard
  Clipboard_Instance,
  Clipboard_SetClipboard,

  // TCanvas
  Canvas_BrushCopy,
  Canvas_CopyRect,
  Canvas_Draw1,
  Canvas_Draw2,
  Canvas_DrawFocusRect,
  Canvas_FillRect,
  Canvas_FrameRect,
  Canvas_StretchDraw,
  Canvas_TextRect1,
  Canvas_TextRect2,

  // TImageList
  ImageList_Draw1,
  ImageList_Draw2,
  ImageList_DrawOverlay1,
  ImageList_DrawOverlay2,
  ImageList_GetIcon1,
  ImageList_GetIcon2,

  // TStyleManager
  StyleManager_IsValidStyle,
  StyleManager_IsValidStyle2,
  StyleManager_LoadFromFile,
  StyleManager_CheckSysClassName,
  StyleManager_TrySetStyle,
  StyleManager_SetStyle1,
  StyleManager_SetStyle2,
  StyleManager_TryLoadFromResource,

  StyleManager_ActiveStyle,
  StyleManager_SystemStyle,
  StyleManager_Enabled,
  StyleManager_IsCustomStyleActive,
  StyleManager_UnRegisterStyle,
  StyleManager_RegisterStyle,
  StyleManager_Style,
  StyleManager_StyleDescriptor,
  StyleManager_StyleNamesOf;