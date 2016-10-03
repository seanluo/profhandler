## Embed profile easily to existing project

### Start / stop profiling by HTTP (For HTTP server programs)

  import github.com/seanluo/profhandler

  Add profhandler.HTTPStart and profhandler.HTTPStop as your handler

  Visit your URL to HTTPStart with parameters: http://xxx/yyy?[PROFILING_MODE]&path=[PATH_FOR_PROFILING_FILE] to start profiling

  > PROFILING_MODE: cpu, mem, block

  > PATH_FOR_PROFILING_FILE: a path on your machine, use "." for current directory

  Visit your URL to HTTPStop to stop profiling

### Start / stop profiling by sending signal to process

  import github.com/seanluo/profhandler

  Add `profhandler.NewSignalHandler()` at the start of your program

  set environment variable PROFILING_MODE to cpu, mem or block

  `kill -SIGUSR1 YOUR_PROCESS_ID` to start profiling

  `kill -SIGUSR2 YOUR_PROCESS_ID` to stop profiling

  Profiling files will be placed in current directory.
