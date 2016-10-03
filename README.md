## Embed profile easily to existing project
### Start / stop profiling by HTTP (For HTTP server programs)
  import github.com/seanluo/profhandler
  Add profhandler.HTTPStart and profhandler.HTTPStop as your handler
  Visit your URL to HTTPStart with parameters: http://xxx/yyy?[PROFILING_MODE]&path=[PATH_FOR_PROFILING_FILE]
  > PROFILING_MODE: cpu, mem, block
  > PATH_FOR_PROFILING_FILE: a path on your machine, use "." for current directory
### Start / stop profiling by sending signal to process
  import github.com/seanluo/profhandler
  Add `profhandler.NewSignalHandler()` at the start of your program
