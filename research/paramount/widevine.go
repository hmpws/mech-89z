package paramount

import (
   "fmt"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "strings"
)

func widevine() {
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "cbsi.live.ott.irdeto.com"
   req.URL.Path = "/widevine/getlicense"
   req.URL.Scheme = "https"
   req.Header["Authorization"] = []string{"Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6IjNkNjg4NGJmLWViMDktNDA1Zi1hOWZjLWU0NGE1NmY3NjZiNiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhbm9ueW1vdXNfVVMiLCJlbnQiOlt7ImJpZCI6IkFsbEFjY2Vzc01haW4iLCJlcGlkIjo3fV0sImlhdCI6MTY0NjQ5NTUzNCwiZXhwIjoxNjQ2NTAyNzM0LCJpc3MiOiJjYnMiLCJhaWQiOiJjYnNpIiwianRpIjoiNDBjY2YxZjQtODkwMi00NjRlLThlMDQtODViYzkyZDcwODJmIn0.YWdeKpytGHA01NxMRsLpd_W7ZEQJ6Fs-Au0lfUgpfr0"}
   val := make(url.Values)
   val["AccountId"] = []string{"cbsi"}
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%q\n", buf)
}

var body = strings.NewReader("\b\x01\x12\x92\x0f\n\x9d\x0e\b\x01\x12\x9a\v\n\xdd\x03\b\x02\x12\x10\xa7ނP\xef\xc1\xfc\x9a\x81\x00+\x17\x06\xa4\xd3\v\x18\x84̋\x91\x06\"\x8e\x020\x82\x01\n\x02\x82\x01\x01\x00\xf9|\xe8\x10@\xb0\x82\x87ʂ\xbb\xa1_\x96\xdfϛPV\u07b2:*\x93Fw\x95\xac+\xe98\xd8n\x95P\f\x96\x90Є\x033nP\u009eT\xba;@aL\x1b\x8e%\U00084408\xe7T\xf3\\\xc0\xd9r~\xf0\xa4MO\xa3v\xed\xb4m\xfa'5v\x8c{\xe0ֱAK\xf7\x0f\x9c2\x19\xe8\xe7*^ddn\x92\xd8\xf2a\xbc\xff\xfa@\xf19\vtE\x1al%o\xdc<\x13h\xe9\x1d\xd3]\xbfS\xa3Eh\x1b\xd1\xd2~&\xbe\xbb\xafc\xb1\xd0R\xbbV\x00\xb3\x9f\xd2\fݑ\xed\xbb;\xa3s\r-\x8a k\x91sf\xd2!\xb5\n\x85.\xfb\xa5\xb6+\xfa\\b\xaa\x9bט\xf9\x03%P\x17U.y\x94\xed>Y\x87\xf1\xca\xce2\xea\\\xa7\x9b\x1b\\\x13\x83\x17D\x87\xa9\x9b\x93\xb9's0\x00GL\xc5\nH\xc3\xff\x0fM\xa5\xe2\xe5\xe2\xd8\xce\xfd(\xa1:\vz%f\r%\xa6\xd6\x00\n\xb1\x9ae\xc2װ\xf6\xde>\x8c\xdf\x02\x03\x01\x00\x01(\xf0\"H\x01R\xaa\x01\b\x01\x10\x00\x1a\x81\x01\x04mC\xb2&\x11\xa6\xfa\xa6\x04q\xe5\xcb\x1d\xb3\x1a)\x9a\xe8\xc8\xf1t\xbf\xae:Iz\xc1\x85e\xee\x8a\u007f\xd9g.d\x87\xfcV}G\x9a(&ai\x8c~\xee\x06q\xd5k\x0eT\xcb4\n\x8a\xf3\xd5>\xc24\x05\x10Zr\xae\x84W\xd4\xd6\xd2\xc3hz\xd2\xe32Y,q\xb8\x9e ̿\xde[_N\xa9\x9c\x16Pm\x89\xcaçyT\xf7\x12)\xb5\x8c\xdc\xec\x18^\u038d\x19)?s\xdc~\xee\x8ez\x14z\x04\x84\xc2\" \xef*\xe2\xe1IH\xe4\xa7\xfe\xe1H\x03\xa1ئ\x9bŷ\xb4;F@;,\xcetB\x9d\f\xf5f1\x12\x80\x02\x99=\xe5w\x15'T\xe3\xca\xe2\xb9b\x1b$U\x84c\xaa2\xadD\xf0Ő\x97\xbe\xa4v\xfb\x8e\x9b\xfa\xa8ٳ\xa7O\xd4@\x82\x95\xc4\xc1\x15\xc9݈Wr\v\x81qyOr\x1a\xcdA\x84(`\x19\x82\xe0\x17\xa4\x98\x17.A5\vbU4\x904\xdf\x04\xe8PסZa\x1a\xd6%\x8f7\xa5~O\xad\xee\bc\xe4i-ܠ\xbb\xc7ơ\xae\xeaZ\\\xec7\x83`\xa6Q3\xceZ\x11=Y=c\rCw\xf6\xb0\u008fA8h\x8a9\xf50h\x81\x8eK\xdbP!\xcf\xf4#\x19\x92\r\x1e\xa9?i!$5\xd2\x12w]aa\x87`\xe3z\xfa^\xe2\x1c\x92\xffE\x04\x14^\xfc\xaah\xf6Z@\x89\x96M|z\x8au\x87@r\xed\xef\xe7H\xa0d\x9d2\x05\x9c\xcb\x12\x15xo\xc2/\xcdk߹\xa1`\xe5\xd5W\xf5\xa1\xf1\xb7|͔\xe36pDآأ\xe6r\t\x9eaq\xa09\xe6;\x820q\xf4s[\xac\x00\x1cu\x1a\xb4\x05\n\xae\x02\b\x01\x12\x10i\xe3蘻,?\xb8\xa3\xb3(\x1d\x84\xf8\x8c\x14\x18\x8eվ\x91\x05\"\x8e\x020\x82\x01\n\x02\x82\x01\x01\x00\xd8\xf6鵉\xf0Q \xe9\xa4>\xd0\xd9N\xa1\xfe\t\x95\x01m\xbd\x1e+@\xa7\xdc\xd1\xc5\u007f\xc3\x05P<\xcf\x13?\x9e\x98\xaf\xceN\xe6\xff\x84\xdcB$&\xa8\xff\xfaO\xe4\xbf-D\xd5\x0f\x14:\xeb\xbc L\xa3\xb4g\xfa%\x1b\xfa`?\xdb%\xe2'\xa8\xaf\xa9\xc3}\n\xefb棔\xd7((\x9fԖe\x1b.\x8cQA\xf2|U\xad\xe69/7\xaa\xd3o7\xce}B\x83[-q~-8t\xfb\xba\xd3\xf3\x14o\xd1x1P\xb7C\xbf\x18\xb9sW\x00tw['\xdd\"\x8c;\x85ˎ\x16]\x9d\xca\xed\x17\xd8厘;\xc63\b\x96u\x89R\xacC\xa3\xb4Б<\xad&e%\xd2G\f\xed\xb9\xdd\x04\xb7\xab\x01\xd2E\x19\xa5\xdcꄘ\xe1\xe3r\xfc\x81\x83\x96\xe2\xc2A\x1d\xcadD,\xcb\xf9\x8by]`\x81\x0e80Z\xfb\\\xe3>\xad߸\xff}x\xab\xf7^A\xb8Ǫ5\x85\xb9Z\xfc\x88\x86\x8an\xaf\x91\xebt\xa1\x02\x03\x01\x00\x01(\xf0\"\x12\x80\x03\"\x8bc0\xeb?\x90\xb3V\xef\xbf\x11\x92\xb6'\xbd\xb6\x97 \xa0\xed\xae\x02\xdeY\xbe\xfd\x02\r{\xeb\x16\\'I\xba>\xa9\x1e\xf5F툕\x95\x95I\xb3y]\x84\x88\x0f\a0\a1\xe3c\xcaͯw\xe0\xb1{R&ܸۗc5\vE\xe7\xe2v\x8e\xb0!\u007ft6\x9b\xb1\b\xe5\xaa!djC\xd8M\x1c\xaf\xe7U\x8f\xda<\xc7K\x89GcVG\xc1U\xa1T\x12]\x9aZws?\xe1\f\xc1,\xdbh\xfe\xdc\xfeP\xac\x99\xc1\xae\xd2s\xf7\x93A;e\xc0\xd2ݑ\x0e\xdeQ\x91\xb0\u007f:\xe0\xa7\xffpA\xaf%\x10A\xa9\xcaI\\\xa7\xc6\xf8\x8bj\x11m@\x00\xfam\x8e\xe8 \xcefD\xcd\x1a\xb9\xc51\x9f\xea\x8a\xfemzʌ*AKѵH\xec\xaf\b\xd1u\x93\x9c\xfa\x94o\n\xb1x`C\x12\xfa\x001̢\xe9q\xbdB\"\x87#\xd8\xe4\xcab\x99\xd5\xdfQ\xd1\x10\xe4p6\x12\xa3b5SH;\x8a\x8f\xd3N\xe3[\x9dr\x01\xe4\xe1:\xbc\xb8'm\xb1\xb3\xc4]\x90\x80\xa58`\x1a\xb8\xb2%\xd2R\xc0\xe1\xf9L\xcdo\x95}\tyŌO\xd0Nm\xba\x19\xf4\x87\xef8\xe0;%S\xc6E\xedM]Sd)\xac\x80)}\xf5\xc3\xf7\b\x00\x00)6\xe9\x18 \xcbǽ\xeb?\"\xf3\xaf\xfea\x8c\x99BG=\x84\x87B\x00Dk\x80F\x02\x11\xb0r\xcc\xe3ku\x10\xec\xd4\xca\xd2\xfd\xc8u\x9a\x85y\xa4\x9b\xfa\xdc\x12\n\xbd\x01Ƽ\xb9\x1a\x16\n\fcompany_name\x12\x06Google\x1a'\n\nmodel_name\x12\x19Android SDK built for x86\x1a\x18\n\x11architecture_name\x12\x03x86\x1a\x1a\n\vdevice_name\x12\vgeneric_x86\x1a$\n\fproduct_name\x12\x14sdk_google_phone_x86\x1aX\n\nbuild_info\x12Jgoogle/sdk_google_phone_x86/generic_x86:7.0/NYC/6696031:userdebug/dev-keys\x1a-\n\tdevice_id\x12 OLrKGvOaFxOkxJEtSvomJKftlXqqMhj\x00\x1a&\n\x14widevine_cdm_version\x12\x0ev4.0.0-android\x1a$\n\x1foem_crypto_security_patch_level\x12\x0102\b\x10\x01 \x00(\v0\x00\x12`\n^\n8\b\x01\x12\x10;输|\x98H1\x84\xb2\x94\x17?\x91R\xaf\" eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU8\x01\x10\x01\x1a 26CCAC944A4BC4000200000000000000\x18\x01 ώ\x8e\x91\x060\x158\xb7\x87\xeb\x9f\x04\x1a\x80\x02]\x12\x1f>T\\,\xfd\xb5W\\6\xc2F&`1v\xe2c]\xe6^\x95\xe3S\xff?^ߟ\xe1\xc6Y,e\xf8p7/KϷ\xcd\xea\xc20\x02,i3!\xf6\xfc;\u05f9\xee\x8eH\xdaQ\xeb\xbbrŋm\x87'\xd6\xf8/|\xb5\xeb\xf7\xd4\tê\xc2z\xa8\xfc\x0f\xae\xacx\t\x1a\x98\xd8㴹\xa3?\b\xe9n3\xfb\xc24\x86\x9f\xa2\x86\xf4\x05\xbc\xb3M\xbcy\xf7\x81\x90\x80\xe4]\xb5\xa8\xe6\xed^p\xb3\xceݔq\x03\xef\xe8\xd7\xee蛳\xa4\x94\xa3:u\x8d\xb7n\xa0\xda5ioV\xe6\xccSd\r\\žHχ\xff\xa3}\x039\x06A{K\xa3\xda\xee+c'\xd3\x13\xce\xd0ܞ\xe5\xe2\xcc\x04\xc3\n\xb2\x0e\x90|'\t\x04\xdcc*9\f\xf0ܒO\xf4f\xa12\x87\x92߶\xf2\xf5\xa9&,\xb85\xe4\u007f\xed\xdf\x12B\xf4|\xee\xef'\t\x0fJ\x1b\x05\x05x\x9e*mn'7BI\x0eB\xba\x04\x83\xc0")
