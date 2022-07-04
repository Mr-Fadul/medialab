# PIPELINES GST-LAUNCH

## Pipeline use plugins interpipes (Ref: https://developer.ridgerun.com/wiki/index.php/GstInterpipe_-_GstInterpipe_Overview)

-> Decode Video VAAPI (vaapih264dec) 

```bash
PIPE_1: dvbsrc name=dvbsrc ! queue ! interpipesink name=input_intp

PIPE_2: interpipesrc name=output_intp listen-to=input_intp is-live=true ! queue ! tsdemux ! h264parse ! queue ! vaapih264dec ! videoscale ! videorate max-rate=30 ! capsfilter caps=video/x-raw,width=640,height=360 ! videoconvert ! deinterlace ! ximagesink

```
-> Decode Audio

```bash
PIPE_1: dvbsrc name=dvbsrc ! queue ! interpipesink name=input_intp

PIPE_2: interpipesrc name=output_intp listen-to=input_intp is-live=true ! queue ! tsdemux ! aacparse ! queue ! avdec_aac_latm ! audioconvert ! audioresample ! pulsesink

```

-> Transcode Video VAAPI for file (vaapih264dec + vaapih264enc) / Transcode Audio

```bash
PIPE_1: dvbsrc name=dvbsrc ! queue ! interpipesink name=input_intp

PIPE_2: interpipesrc name=output_intp listen-to=input_intp is-live=true ! queue ! tsdemux name=demux  ! h264parse ! vaapih264dec ! queue ! videoscale ! videorate max-rate=15 ! capsfilter caps="video/x-raw,width=640,height=360" ! videoconvert ! deinterlace ! vaapih264enc bitrate=1 ! h264parse ! queue ! mp4mux name=mux ! filesink name=filesink location=/home/edmilson/Videos/transcode_audio_video.mp4 demux. ! aacparse ! avdec_aac_latm ! queue ! audioconvert ! audioresample ! avenc_aac ! mux.

```

### Transcode Video VAAPI for file (vaapih264dec + vaapih264enc) / Transcode Audio
GST_DEBUG=1 gst-launch-1.0 dvbsrc frequency=497142857 inversion="off" ! tsdemux name=demux  ! h264parse ! vaapih264dec ! queue ! videoscale ! videorate max-rate=15 ! capsfilter caps="video/x-raw,width=640,height=360" ! videoconvert ! deinterlace ! vaapih264enc bitrate=1 ! h264parse ! queue ! mp4mux name=mux ! filesink location=/home/edmilson/Videos/transcode_audio_video.mkv -e demux. ! aacparse ! avdec_aac_latm ! queue ! audioconvert ! audioresample ! avenc_aac ! mux.

### Transcode Video VAAPI for file (vaapih264dec + x264enc - H.264 Software Video Encoder) / Transcode Audio
GST_DEBUG=1 gst-launch-1.0 dvbsrc frequency=497142857 inversion="off" ! tsdemux name=demux ! h264parse ! vaapih264dec ! queue ! videoscale ! capsfilter caps=video/x-raw,width=640,height=360 ! x264enc name=videoenc bitrate=600 ! queue ! mp4mux name=mux ! filesink name=filesink location=~/Videos/transcode_video_$(date '+%Y%m%d_%H%M%S').mp4 -e demux. ! aacparse ! avdec_aac_latm ! queue ! audioconvert ! audioresample ! avenc_aac ! queue ! mux.

### Transcode Video VAAPI for file  (vaapih264dec + vaapih264enc)
GST_DEBUG=1 gst-launch-1.0 dvbsrc frequency=497142857 inversion="off" ! tsdemux ! h264parse ! vaapih264dec ! queue ! videoscale ! video/x-raw,width=640,height=360 ! vaapih264enc bitrate=600 ! h264parse ! mp4mux ! filesink location=~/Videos/transcode_video_$(date '+%Y%m%d_%H%M%S').mkv -e 

### Transcode Audio for file 
GST_DEBUG=1 gst-launch-1.0 dvbsrc frequency=497142857 inversion="off" ! tsdemux ! aacparse ! queue ! avdec_aac_latm ! audioconvert ! audioresample ! avenc_aac ! mp4mux ! filesink location=.~/Videos/transcode_audio_$(date '+%Y%m%d_%H%M%S').mkv -e 

### Decode Video VAAPI (vaapih264dec) 
GST_DEBUG=1 gst-launch-1.0 dvbsrc frequency=497142857 inversion="off" ! tsdemux !  h264parse ! queue ! vaapih264dec ! videoscale ! videorate max-rate=30 ! capsfilter caps=video/x-raw,width=640,height=360 ! videoconvert ! deinterlace ! ximagesink

### Decode Audio
GST_DEBUG=1 gst-launch-1.0 dvbsrc frequency=497142857 inversion="off" ! tsdemux ! aacparse ! queue ! avdec_aac_latm ! audioconvert ! audioresample ! pulsesink