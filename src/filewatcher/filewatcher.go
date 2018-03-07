package filewatcher

import (
	"log"

	"github.com/howeyc/fsnotify"
)

/**
* Интерфейс для "откладывания" закрытия
 */
type Closeable interface {
	Close()
}

/**
* Объект для "откладывания" закрытия
 */
type CloseableWrapper struct {
	watcher *fsnotify.Watcher
}

func (wrapper *CloseableWrapper) Close() {
	wrapper.watcher.Close()
}

/**
*Наблюдать за каталогом
 */
func Watch(folder string, c chan string) *CloseableWrapper {
	watcher, err := fsnotify.NewWatcher()
	handleError(err)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if ev.IsCreate() {
					c <- ev.Name
				} else if ev.IsDelete() {
					c <- "Delete: " + ev.Name
				}
			case err := <-watcher.Error:
				handleError(err)
			}
		}
	}()

	err = watcher.Watch("/home/alex/tmp")
	handleError(err)

	closeableWrapper := CloseableWrapper{watcher: watcher}
	return &closeableWrapper
}

/**
*Обработать ошибку
 */
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
