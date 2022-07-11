package main

import (
   "github.com/89z/mech/paramount"
   "github.com/89z/rosso/dash"
   "github.com/89z/rosso/hls"
   "strings"
)

func (f flags) DASH(preview *paramount.Preview) error {
   address, err := paramount.New_Media(f.guid).DASH()
   if err != nil {
      return err
   }
   f.Poster, err = paramount.New_Session(f.guid)
   if err != nil {
      return err
   }
   f.Base = preview.Base()
   reps, err := f.Stream.DASH(address.String())
   if err != nil {
      return err
   }
   audio := reps.Filter(func(r dash.Representation) bool {
      if r.MimeType != "audio/mp4" {
         return false
      }
      if r.Role() != "" {
         return false
      }
      return true
   })
   index := audio.Index(func(a, b dash.Representation) bool {
      if !strings.Contains(b.Codecs, f.codecs) {
         return false
      }
      if b.Adaptation.Lang != f.lang {
         return false
      }
      return true
   })
   if err := f.DASH_Get(audio, index); err != nil {
      return err
   }
   video := reps.Video()
   return f.DASH_Get(video, video.Bandwidth(f.bandwidth))
}

func (f flags) HLS(preview *paramount.Preview) error {
   address, err := paramount.New_Media(f.guid).HLS()
   if err != nil {
      return err
   }
   f.Base = preview.Base()
   master, err := f.Stream.HLS(address.String())
   if err != nil {
      return err
   }
   streams := master.Streams.Filter(func(s hls.Stream) bool {
      return s.Resolution != ""
   })
   return f.HLS_Streams(streams, streams.Bandwidth(f.bandwidth))
}
