// Docker Volume GIT plugin allows to mount git repository in container.
// Copyright (C) 2017 Kassisol inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	h, err := newHandlerFromVolumeDriver("/var/lib/docker")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(h.ServeUnix("gitvol", 1))
}
