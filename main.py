#! /usr/bin/env python
# -*- coding: utf-8 -*-
                             
import re, urllib, urllib2
                             
class XiamiDownload(object):
                             
    """虾米音乐下载"""
                             
    def __init__(self, url_song):
                             
        """ 初始化，得到请求xml和加密的下载地址 """
                             
        self.url_song = url_song      
        self.url_xml = self.__get_xml()
        self.info = self. __get_info()
        self.url_location = self.info
        #self.lyc = self.info[1]
        #self.pic = self.info[2]
                                 
    def __get_xml(self):
                             
        """ 得到请求的 xml 地址 """
                             
        return 'http://www.xiami.com/song/playlist/id/%s/object_name/default/object_id/0' % re.search('\d+', self.url_song).group()
                             
    def __get_info(self):
                             
        """ 伪装浏览器请求，处理xml，得到 加密的 location """
                             
        headers = {
            'User-Agent':'Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.6) Gecko/20091201 Firefox/3.5.6'
        }
        req = urllib2.Request(
            url = self.url_xml,
            headers = headers
        )
        xml = urllib2.urlopen(req).read()
        location = re.search(r'(<location>)(.+)(<\/location>)', xml).group(2)
        #lyc = re.search(r'(<lyric>)(.+)(<\/lyric>)', xml).group(2)
        #pic = re.search(r'(<pic>)(.+)(<\/pic>)', xml).group(2)
        #return (location, lyc, pic)
        return location
                             
    def get_url(self):
                             
        """ 解密 location 获得真正的下载地址 """
                     
        strlen = len(self.url_location[1:])
        rows = int(self.url_location[0])
        cols = strlen / rows
        right_rows = strlen % rows
        new_str = self.url_location[1:]
        url_true = ''
                             
        for i in xrange(strlen):
            x = i % rows
            y = i / rows
            p = 0
            if x <= right_rows:
                p = x * (cols + 1) + y
            else:
                p = right_rows * (cols + 1) + (x - right_rows) * cols + y
            url_true += new_str[p]
        return urllib2.unquote(url_true).replace('^', '0')
                             
if __name__ == '__main__':
                             
	url = raw_input('please enter the url of the song: ')
	#url = "http://www.xiami.com/song/1769629568"
	xi = XiamiDownload(url) 
	url_download = xi.get_url()
	#url_pic = xi.pic
	#url_lyc = xi.lyc
    
	urllib.urlretrieve(url_download, 'new.mp3')
	#urllib.urlretrieve(url_pic, r'new.jpg')
	#urllib.urlretrieve(url_lyc, r'new.lyc')
