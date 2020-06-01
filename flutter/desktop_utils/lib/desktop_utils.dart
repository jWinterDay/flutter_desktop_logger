import 'dart:io';
import 'package:flutter/services.dart';

// method channel names
const MethodChannel _loggerChannel = MethodChannel('github.com/jWinterDay/desktop_logger');

class DesktopUtils {
  /// desktop log
  /// [path] may be relative or absolute
  static void logToFile(
    String text, {
    String path = 'desktop_log.txt',
    String prefix = '[DEBUG] ',
  }) {
    if (Platform.isAndroid || Platform.isIOS) {
      // print('not supported platform: ${Platform.operatingSystem}');
      return;
    }

    _loggerChannel.invokeMethod(
      'log',
      <String, dynamic>{
        'path': path,
        'prefix': prefix,
        'text': text,
      },
    );
  }
}
