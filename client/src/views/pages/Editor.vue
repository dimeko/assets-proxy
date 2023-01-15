<template>
  <div class="editor_container">
    <div class="editor">
      <div class="editor_header">
        <router-link :to="{ name: 'dashboard' }"
          ><div>Dashboard</div></router-link
        >
      </div>
      <div class="editor_content">
        <div class="editor_file_management">
          <div class="editor_directories">
            <div class="editor_search_file_title">Search file</div>
            <div class="editor_file_search">
              <label for="file">File name</label
              ><input type="text" v-model="fileName" />
            </div>

            <div class="editor_button">
              <button @click="fetchFile">Search file</button>
            </div>

            <div class="editor_search_file_title" style="margin-top: 18px;">Create file</div>
            <div class="editor_file_search">
              <label for="file">File name</label
              ><input type="text" v-model="createFileName" />
            </div>

            <div class="editor_file_search" style="margin-top: 10px;">
              <label for="file">File subfolder</label
              ><input type="text" v-model="createSubfolder" />
            </div>

            <div class="editor_file_search" style="margin-top: 10px;">
              <label for="file">Images subfolder</label
              ><input type="text" v-model="createSubfolderImages" />
            </div>

            <div class="editor_button">
              <button @click="createFile">Create file</button>
            </div>
          </div>
          <div class="editor_image_upload">
            <form enctype="multipart/form-data">
              <div class="editor_upload_image_title">Upload image</div>
              <div class="editor_upload_image_directory">
                <label for="directory">Directory</label>
                <select
                  class="form-control"
                  :required="true"
                  v-model="imageDirectory"
                  @change="fetchSubdirImages"
                >
                  <option selected>Choose directory</option>
                  <option
                    v-for="directory in directories"
                    v-bind:value="directory"
                    v-bind:key="directory"
                  >
                    {{ directory }}
                  </option>
                </select>
              </div>
              <div class="dropbox">
                <input
                  type="file"
                  :name="uploadFieldName"
                  @change="
                    onImageChange($event.target.name, $event.target.files)
                  "
                  accept="image/*"
                  class="input-file"
                />
                <img
                  v-if="imagePreviewUrl"
                  class="dropbox_image_preview"
                  :src="imagePreviewUrl"
                  alt="uploading image"
                />
                <p v-if="!imagePreviewUrl">
                  Drag your file(s) here to begin<br />
                  or click to browse
                </p>
              </div>
            </form>
            <div class="editor_image_upload_button">
              <button @click="uploadImage">Upload image</button>
            </div>
          </div>
          <div
            v-if="directoryImages && directoryImages.length > 0"
            class="editor_directory_images"
          >
            <div class="editor_directory_images_title">
              Currenct directory images
            </div>
            <div class="editor_directory_image_content">
              <div
                class="editor_directory_image_wrapper"
                v-for="image in directoryImages"
                :key="image"
              >
                <p class="editor_directory_image_name">Image name: {{ image }}</p>
                <img
                  class="editor_directory_image"
                  :src="`${baseApiUrl}image/${imageDirectory}/${image}`"
                />
              </div>
            </div>
          </div>
        </div>

        <div class="editor_editor">
          <vue-json-editor
            v-model="json"
            :show-btns="true"
            :expandedOnStart="true"
            @json-change="onJsonChange"
            @json-save="editDbFile"
          ></vue-json-editor>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import vueJsonEditor from "vue-json-editor";
import axios from "axios";

export default {
  data() {
    return {
      fileName: "",
      imageName: "",
      imageDirectory: "",
      uploadFieldName: "image_file",
      imagePreviewUrl: "",
      imageFile: null,
      json: {
        msg: "demo of jsoneditor",
      },
      directories: [],
      directoryImages: [],
      baseApiUrl: "/api/server/",
      createFileName: "",
      createSubfolder: "",
      createSubfolderImages: ""
    };
  },
  components: {
    vueJsonEditor,
  },
  mounted() {
    this.mapImageDirectory();
  },
  methods: {
    onImageChange(name, files) {
      this.imageFile = files[0];
      this.imagePreviewUrl = URL.createObjectURL(files[0]);
    },
    async fetchFile() {
      await axios
        .get(`${this.baseApiUrl}get-file?file_path=${this.fileName}`)
        .then((response) => {
          this.json = response.data.body;
        })
        .catch(() => {
          this.$root.$emit("notify", {
            title: "Error",
            description: "File not found!",
            bodyClass: "error_toast",
          });
        });
    },
    async mapImageDirectory() {
      const url = `${this.baseApiUrl}map-image-directory`;
      await axios.get(url).then((response) => {
        this.directories = response.data.body;
      });
    },
    async uploadImage() {
      const url = `${this.baseApiUrl}upload-image`;
      let formData = new FormData();
      formData.append("image_file", this.imageFile);
      formData.append("category", this.imageDirectory);
      await axios
        .post(url, formData)
        .then(() => {
          this.$root.$emit("notify", {
            title: "Success",
            description: "The image was successfully uploaded",
            bodyClass: "success_toast",
          });
        })
        .catch(() => {
          this.$root.$emit("notify", {
            title: "Error",
            description: "Error while uploading the image!",
            bodyClass: "error_toast",
          });
        });
    },
    async editDbFile() {
      const url = `${this.baseApiUrl}edit-db-file`;
      await axios
        .post(url, {
          file_path: this.fileName,
          data: this.json,
        })
        .then(() => {
          this.$root.$emit("notify", {
            title: "Success",
            description: "The file was successfully updated",
            bodyClass: "success_toast",
          });
        })
        .catch(() => {
          this.$root.$emit("notify", {
            title: "Error",
            description: "Error while updating the file!",
            bodyClass: "error_toast",
          });
        });
    },
    async fetchSubdirImages() {
      const url = `${this.baseApiUrl}map-image-directory?sub_dir=${this.imageDirectory}`;
      await axios.get(url).then((response) => {
        this.directoryImages = response.data.body;
      });
    },
    async createFile() {
      const url = `${this.baseApiUrl}create-file`;
      this.createFileName = this.createFileName[0] !== "/" ? "/" + this.createFileName : this.createFileName
      this.createSubfolder = this.createSubfolder[0] !== "/" ? "/" + this.createSubfolder : this.createSubfolder
      this.createSubfolderImages = this.createSubfolderImages[0] !== "/" ? "/" + this.createSubfolderImages : this.createSubfolderImages

      await axios.post(url, {
        file_name: this.createFileName,
        subfolder: this.createSubfolder,
        img_subfolder: this.createSubfolderImages,
        data: {}
      }).then(() => {
          this.$root.$emit("notify", {
            title: "Success",
            description: "The file was successfully created",
            bodyClass: "success_toast",
          });
        })
        .catch(() => {
          this.$root.$emit("notify", {
            title: "Error",
            description: "Error while creating file!",
            bodyClass: "error_toast",
          });
        });
    }
  },
};
</script>
