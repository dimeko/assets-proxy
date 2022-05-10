<template>
  <div class="editor_container">
    <div class="editor">
      <div class="editor_header">
        <router-link  :to="{ name: 'dashboard' }"
          ><div>Dashboard</div></router-link
        >
      </div>
      <div  class="editor_content">
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
const BASE_API_URL = "/api/server/";
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
    };
  },
  components: {
    vueJsonEditor,
  },
  mounted() {
    this.mapImageDirectory();
  },
  methods: {
    onJsonChange(value) {
      console.log("value:", value);
    },
    onImageChange(name, files) {
      console.log(files);
      this.imageFile = files[0];
      this.imagePreviewUrl = URL.createObjectURL(files[0]);
    },
    async fetchFile() {
      const response = await axios.get(
        `${BASE_API_URL}get-file?file_path=${this.fileName}`
      );
      this.json = response.data.body;
    },
    async mapImageDirectory() {
      const url = `${BASE_API_URL}map-image-directory`;
      const response = await axios.get(url);

      this.directories = response.data.body;
    },
    async uploadImage() {
      const url = `${BASE_API_URL}upload-image`;
      let formData = new FormData();
      formData.append("image_file", this.imageFile);
      formData.append("category", this.imageDirectory);
      console.log(this.imageDirectory, this.imageName);
      const response = await axios.post(url, formData);
      console.log(response)
    },
    async editDbFile() {
      const url = `${BASE_API_URL}edit-db-file`;
      const response = await axios.post(url, {
        file_path: this.fileName,
        data: this.json,
      });
      console.log(response)
    },
  },
};
</script>